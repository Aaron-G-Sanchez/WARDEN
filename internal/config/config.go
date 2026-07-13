package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	WebhookURL string `json:"webhook_url"`
	Token      string `json:"token"`
}

// TODO: Update function to create warden directory if it does not exist.
func GetConfig() (*Config, error) {
	// TODO: Create directory if it does not exist.
	// Look for the file
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := path.Join(home, ".config", "warden", "config.json")

	// TODO: Read file
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
