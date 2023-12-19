package setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func CreateServiceFile() error {
	serviceContent := `[Unit]
Description=puller

[Service]
Environment=PORT=8080
Environment=GO_ENV=production
Environment=GIN_MODE=release
Type=simple
User=%s
Restart=always
RestartSec=5s
ExecStart=%s/puller -s
WorkingDirectory=%s

[Install]
WantedBy=multi-user.target
`

	user, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	serviceContent = fmt.Sprintf(serviceContent, user.Username, currentDir, currentDir)
	serviceFileName := "puller.service"
	serviceFilePath := filepath.Join("/lib/systemd/system", serviceFileName)

	err = os.WriteFile(serviceFilePath, []byte(serviceContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing service file: %w", err)
	}

	cmd := exec.Command("sudo", "systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running systemctl daemon-reload: %w", err)
	}

	// Allow port 8080 through the firewall
	ufwCmd := exec.Command("sudo", "ufw", "allow", "8080")
	ufwErr := ufwCmd.Run()
	if ufwErr != nil {
		return fmt.Errorf("error allowing port 8080 through the firewall: %w", ufwErr)
	}

	// Start the puller service
	startCmd := exec.Command("sudo", "service", "puller", "start")
	startErr := startCmd.Run()
	if startErr != nil {
		return fmt.Errorf("error starting puller service: %w", startErr)
	}

	return nil
}

func GenerateConfigFile() error {
	const configFile = "config.json"
	config := map[string]interface{}{
		"key":            "",
		"email":          "email@example.com",
		"email_password": "password",
		"smtp_host":      "smtp.pramixit.com",
		"smtp_port":      "587",
		"repos": []map[string]interface{}{
			{
				"name":   "repo-name",
				"path":   "local-path",
				"branch": "branch-name",
				"events": []map[string]interface{}{{"type": "push", "commands": []string{"git pull"}}},
			},
		},
	}

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 25)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	config["key"] = string(b)

	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}
	err = ioutil.WriteFile(configFile, jsonData, 0666)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func OutputBanner() {
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)

	}
	for _, addr := range address {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("\n\nUSE THIS URL AS YOUR GITLAB WEBHOOK URL: http://" + ipNet.IP.String() + ":8080/puller\n\n")
			}
		}
	}
}
