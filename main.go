package main

import (
	"log"
	"os"

	"orltom.dev/golang-http-example/cmd"
)

func main() {
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Could not start web application. %v", err)
		os.Exit(1)
	}
}
