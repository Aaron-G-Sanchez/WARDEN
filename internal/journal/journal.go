package journal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Log struct {
	Timestamp       string `json:"__REALTIME_TIMESTAMP"`
	SystemdUnit     string `json:"_SYSTEMD_UNIT"`
	Message         string `json:"MESSAGE"`
	SyslogTimestamp string `json:"SYSLOG_TIMESTAMP"`
}

func Watch(reader io.Reader) {
	decoder := json.NewDecoder(reader)
	for decoder.More() {

		var logEntry *Log

		if err := decoder.Decode(&logEntry); err != nil {
			log.Printf("Failed to decode log entry: %v\n", err)
		}

		// TODO: Evaluate line.
		// 1) Filter logs for any log before command was ran.
		// 2) Call the notifier when specific connection requests are initiated.
		fmt.Printf("%+v \n", *logEntry)

	}

}

func evaluate(logEntry *Log) {
	// 1) Skip entries older than the program start time
	// 2) Evalutate the logEntry.Message for connection patterns
	// 		and call the notifier if matched.
}
