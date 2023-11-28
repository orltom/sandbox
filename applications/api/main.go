package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"orltom.dev/golang-http-example/cmd"
	"orltom.dev/golang-http-example/internal/repository"
	"orltom.dev/golang-http-example/internal/resources"
	"orltom.dev/golang-http-example/internal/setup"
)

func main() {
	db, err := setup.OpenDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection. %v", err)
		os.Exit(1)
	}
	repository := repository.NewPostgresJokeRepository(db)
	rest := resources.NewJokeRestResource(repository)
	err = cmd.Start(rest)
	if err != nil {
		log.Fatalf("Could not start web application. %v", err)
		os.Exit(1)
	}
}
