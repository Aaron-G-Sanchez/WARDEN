package journal

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Watch() {
	cmd := exec.Command("journalctl", "-t", "sshd-session", "-f")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to create stdout pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start journalctl command: %v", err)
	}
	defer cmd.Process.Kill()

	fmt.Printf("Starting realtime logger")

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		fmt.Println("LINE:::", string(line))

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading stream: %v", err)
		}
	}

	cmd.Wait()
}
