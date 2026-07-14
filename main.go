package main

import (
	"cycle/warden/internal/config"
	"cycle/warden/internal/journal"
	"fmt"
	"log"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Printf("%+v\n", *config)

	// JOURNAL LOGGING TEST
	journal.Watch()
}
