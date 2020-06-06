package main

import (
	"fmt"
	"log"
)

var config appConfig

func init() {
}

func main() {
	// Load Application Configuration
	config, err := ReadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
