package main

import (
	"testing"
)

// Tests whether you can open the right json file
func TestReadConfig(t *testing.T) {
	// Should open file
	if _, err := ReadConfig("config.json"); err != nil {
		t.Errorf("config.go failed to open/read valid config file: %v\n", err)
	}
	// Should not open file (file does not exist)
	if _, err := ReadConfig("config_test_wrong.json"); err == nil {
		t.Errorf("config.go read an invalid config file: %v\n", err)
	}
}
