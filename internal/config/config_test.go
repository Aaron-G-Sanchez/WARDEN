package config

import (
	"encoding/json"
	"os"
	"path"
	"testing"
)

func TestGetConfig_WithConfigFile(t *testing.T) {
	mockConfig := Config{
		WebhookURL: "test.com",
		Token:      "test_token",
	}

	setup(t, mockConfig)

	got, err := GetConfig()
	if err != nil {
		t.Fatalf("Error reading config file: %v", err)
	}

	if *got != mockConfig {
		t.Errorf("Expected: %v, got: %v", *got, mockConfig)
	}
}

func setup(t *testing.T, mockConfig Config) {
	tmpDir := t.TempDir()
	testFile := path.Join(tmpDir, ".config", "warden", "config.json")

	content, err := json.Marshal(mockConfig)
	if err != nil {
		t.Fatalf("Error creating config: %v", err)
	}

	err = os.MkdirAll(path.Dir(testFile), 0750)
	if err != nil {
		t.Fatalf("Error creating config directories: %v", err)
	}

	err = os.WriteFile(testFile, content, 0644)
	if err != nil {
		t.Fatalf("Error writing config: %v", err)
	}

	t.Setenv("HOME", tmpDir)
}
