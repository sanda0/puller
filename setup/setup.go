package setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running systemctl daemon-reload: %w", err)
	}

	return nil
}

func GenerateConfigFile() error {
	const configFile = "config.json"
	config := map[string]interface{}{
		"key": "",
		"repos": []map[string]interface{}{
			{
				"name":          "<repo-name>",
				"path":          "<local-path>",
				"branch":        "refs/heads/<branch-name>",
				"events":        []map[string]interface{}{{"type": "push", "commands": []string{"git pull"}}},
				"notifications": []map[string]interface{}{{"type": "", "to": []string{""}}},
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
