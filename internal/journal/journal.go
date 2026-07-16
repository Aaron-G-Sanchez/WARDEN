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

	// TODO: Remove after testing.
	fmt.Printf("Starting realtime logger \n")
	lineNum := 0

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		// TODO: Evaluate line.
		fmt.Printf("LINE::%v:: %v\n", lineNum, string(line))

		lineNum++
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading stream: %v", err)
		}
	}

	cmd.Wait()
}
