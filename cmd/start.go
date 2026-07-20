package cmd

import (
	"cycle/warden/internal/journal"
	"fmt"
	"log"
	"os/exec"
)

func RunStart() {
	cmd := exec.Command("journalctl", "-t", "sshd-session", "-o", "json", "-f")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to create stdout pipe: %v", err)
	}

	fmt.Printf("Starting logger \n")

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start journalctl command: %v", err)
	}
	defer cmd.Process.Kill()

	journal.Watch(stdout)
	cmd.Wait()
}
