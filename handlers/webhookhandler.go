package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sanda0/puller/stucts"
	"gopkg.in/gomail.v2"
)

func findRepoByName(repoName string) (*stucts.Repo, error) {
	var config stucts.Config

	// Read the config.json file
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the JSON into the Config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Iterate through the repositories to find the matching one
	for _, repo := range config.Repos {
		if repo.Name == repoName {
			return &repo, nil
		}
	}

	return nil, fmt.Errorf("repository not found")
}

func getKeyFromConfig() (string, error) {
	var config stucts.Config

	// Read the config.json file
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return "", fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the JSON into the Config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling config: %w", err)
	}

	return config.Key, nil
}

func getConfig() (*stucts.Config, error) {
	var config stucts.Config

	// Read the config.json file
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the JSON into the Config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

func HandleWebHook(c *gin.Context) {
	secretToken := c.GetHeader("X-Gitlab-Token")
	key, err := getKeyFromConfig()
	if err != nil {
		return
	}
	if secretToken != key {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	payload := stucts.GitLabWebhookPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println("err: ", err)
	}
	repo, err := findRepoByName(payload.Project.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if payload.ObjectKind == "push" {
		refParts := strings.Split(payload.Ref, "/")
		refSuffix := refParts[len(refParts)-1]
		if refSuffix == repo.Branch {
			go runCmd(payload.ObjectKind, repo.Path, repo.Events, payload.UserEmail)
		}
	} else if payload.ObjectKind == "merge_request" {
		if payload.ObjectAttributes.TargetBranch == repo.Branch {

			go runCmd(payload.ObjectKind, repo.Path, repo.Events, payload.User.Email)
		}
	}

	c.JSON(http.StatusOK, "OK")

}

func runCmd(event string, path string, events []stucts.Event, userEmail string) {
	exec.Command("sh", "-c", "git config --global --add safe.directory "+path).Run()
	for _, e := range events {
		if event == e.Type {
			cmd := "cd " + path
			for _, c := range e.Commands {
				cmd = cmd + " && " + c
			}
			fmt.Println(cmd)
			// Execute the cmd string in the shell
			output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
			if err != nil {
				writeLogFile("ERROR", fmt.Sprintf("Failed to execute command: %s\nOutput: %s\nError: %v\n", cmd, string(output), err))
				SendEmail(userEmail, "Failed to execute command", fmt.Sprintf("Failed to execute command: %s\nOutput: %s\nError: %v\n", cmd, string(output), err))
			} else {
				writeLogFile("INFO", fmt.Sprintf("Command executed successfully: %s\nOutput: %s\n", cmd, string(output)))
				SendEmail(userEmail, "Command executed successfully", fmt.Sprintf("Command executed successfully: %s\nOutput: %s\n", cmd, string(output)))
			}
		}
	}
}

func writeLogFile(messageType, message string) error {
	logFile := "app.log"
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("error opening log file: %w", err)
	}
	defer file.Close()

	var logger *log.Logger
	if messageType == "INFO" {
		logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else if messageType == "ERROR" {
		logger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		return fmt.Errorf("invalid message type: %s", messageType)
	}

	logger.Println(message)

	return nil
}

func SendEmail(To string, Subject string, Body string) error {
	config, err := getConfig()
	if err != nil {
		log.Println("Error getting email configuration:", err)
		return err
	}

	message := "To: " + To + "\r\n" +
		"Subject: " + Subject + "\r\n" +
		"\r\n" +
		Body

	m := gomail.NewMessage()
	m.SetHeader("From", config.Email)
	m.SetHeader("To", To)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/plain", message)

	port, err := strconv.Atoi(config.SMTPPort)
	if err != nil {
		log.Printf("Error converting SMTPPort to int: %s\n", err)
		return err
	}
	d := gomail.NewDialer(config.SMTPHost, port, config.Email, config.EmailPassword)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		writeLogFile("ERROR", fmt.Sprintf("Error sending email: %s\n", err))
		return err
	}

	return nil
}
