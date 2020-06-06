#!/bin/bash

### Test go files ###
go test
###

### Runs go vet ###
go vet -json | tee ./govet-output.json
###