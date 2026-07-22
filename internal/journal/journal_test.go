package journal

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

var MockInput = `
	{
	"__SEQNUM" : "1111",
	"_COMM" : "sshd-session",
	"_TRANSPORT" : "syslog",
	"_MACHINE_ID" : "test123",
	"_HOSTNAME" : "test-station",
	"_SYSTEMD_INVOCATION_ID" : "abc123",
	"__REALTIME_TIMESTAMP" : "1784592008629207",
	"PRIORITY" : "6",
	"_SYSTEMD_SLICE" : "system.slice",
	"SYSLOG_PID" : "111",
	"__MONOTONIC_TIMESTAMP" : "342797466",
	"SYSLOG_TIMESTAMP" : "Jul 20 19:00:08 ",
	"SYSLOG_IDENTIFIER" : "sshd-session",
	"_BOOT_ID" : "bootid123",
	"SYSLOG_FACILITY" : "10",
	"_RUNTIME_SCOPE" : "system",
	"_EXE" : "/usr/lib/openssh/sshd-session",
	"_GID" : "0",
	"_SYSTEMD_CGROUP" : "/system.slice/ssh.service",
	"_PID" : "111",
	"_CAP_EFFECTIVE" : "1ffffffffff",
	"MESSAGE" : "pam_unix(sshd:session): session opened for user test(uid=1000) by test(uid=0)",
	"_UID" : "0",
	"_SYSTEMD_UNIT" : "ssh.service",
	"_SOURCE_REALTIME_TIMESTAMP" : "1784592008629149"
	}`

// TODO: IMPLEMENT.
func TestWatch(t *testing.T) {
	mockOutput := &Log{
		Timestamp:       "1784592008629207",
		SystemdUnit:     "ssh.service",
		Message:         "pam_unix(sshd:session): session opened for user test(uid=1000) by test(uid=0)",
		SyslogTimestamp: "Jul 20 19:00:08 ",
	}

	want := fmt.Sprintf("%+v\n", *mockOutput)

	var buf bytes.Buffer
	log.SetFlags(0)
	log.SetOutput(&buf)

	testReader := strings.NewReader(MockInput)

	Watch(testReader)

	got := buf.String()

	if got != want {
		t.Errorf("got log %s; want log %s", got, want)
	}
}
