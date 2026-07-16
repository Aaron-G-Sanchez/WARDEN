package cmd

import (
	"errors"
	"fmt"
	"os"
)

func Execute() error {

	if len(os.Args) < 2 {
		showHelp()
		return errors.New("No command provided")
	}

	subCmd := os.Args[1]

	switch subCmd {
	case "start":
		RunStart()
	case "help":
		showHelp()
	}

	return nil
}

func showHelp() {
	fmt.Println("Usage: warden <command> [arguments]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("	start	-	Start the log monitor")
	fmt.Println("	help	-	List available commands")
}
