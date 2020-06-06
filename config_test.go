package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// Creates Test File
func createTestFile(tf string, content []byte) {
	if err := ioutil.WriteFile(tf, content, 0444); err != nil {
		panic(err)
	}
}

// Removes Test Files
func cleanUp(tf string) {
	if err := os.Remove(tf); err != nil {
		panic(err)
	}
}

// Tests whether you can open the right json file
func TestReadConfig(t *testing.T) {
	var tfPaths []string

	// Test 1: Should open and read file
	tf := "./tmp/config.json"
	json := []byte(`{"name":"Template-Config",
					"description":"Simple Application with Configurations given in a JSON named config.json",
					"serverAddress": "127.0.0.1:8080"
					}`)
	createTestFile(tf, json)
	defer cleanUp(tf)
	tfPaths = append(tfPaths, tf)
	_, err := ReadConfig(tf)
	if err != nil {
		t.Errorf("config.go failed to open/read valid config file: %v\n", err)
	}

	// Test 2: Should not open file (file does not exist)
	tf = "./tmp/config_not_exist.json"
	tfPaths = append(tfPaths, tf)
	if _, err := ReadConfig(tf); err == nil {
		t.Errorf("config.go claims to have opened a non-existent file.\n")
	}

	// Test 3: Should not read file (file has an empty JSON)
	tf = "./tmp/config_wrong.json"
	json = []byte(`{}`)
	createTestFile(tf, json)
	defer cleanUp(tf)
	tfPaths = append(tfPaths, tf)
	if _, err := ReadConfig(tf); err == nil {
		t.Errorf("config.go read invalid config file: %s\n", json)
	}
}
