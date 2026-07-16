package main

import (
	"cycle/warden/cmd"
	"cycle/warden/internal/config"
	"log"
)

func main() {
	_, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
