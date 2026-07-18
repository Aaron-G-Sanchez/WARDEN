package journal

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type Log struct {
	Timestamp       string `json:"__REALTIME_TIMESTAMP"`
	SystemdUnit     string `json:"_SYSTEMD_UNIT"`
	Message         string `json:"MESSAGE"`
	SyslogTimestamp string `json:"SYSLOG_TIMESTAMP"`
}

func Watch() {
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

	decoder := json.NewDecoder(stdout)

	for decoder.More() {

		var logEntry *Log

		if err := decoder.Decode(&logEntry); err != nil {
			log.Printf("Failed to decode log entry: %v\n", err)
		}

		fmt.Printf("%+v \n", *logEntry)

	}

	cmd.Wait()
}
