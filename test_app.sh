#!/bin/bash

# Create files for config_test.go testing
CONFIG_FILE_WRONG_NAME="./config_test_wrong.json"
CONFIG_FILE_WRONG_JSON="{}"

touch "$CONFIG_FILE_WRONG_NAME"
echo "$CONFIG_FILE_WRONG_JSON" > $CONFIG_FILE_WRONG_NAME

go test ./...

# Remove temporary files used by config_test.go
rm "$CONFIG_FILE_WRONG_NAME"

go vet -json | tee ./govet-output.json