package main

import (
	"fmt"
	"log"
)

var config appConfig

func init() {
}

func main() {
	config, err := ReadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
