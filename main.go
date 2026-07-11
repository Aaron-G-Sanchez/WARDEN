package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

type Config struct {
	WebhookURL string `json:"webhook_url"`
	Token      string `json:"token"`
}

func main() {
	// TODO: get config data
	config := getConfig()

	fmt.Printf("%+v\n", config)
}

func getConfig() Config {
	// TODO: Create directory if it does not exist.
	// Look for the file
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home dir: %v", err)
	}
	path := path.Join(home, ".config", "warden", "config.json")

	// TODO: Read file
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return config
}
