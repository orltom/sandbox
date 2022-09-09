package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"orltom.dev/golang-http-example/cmd"
	"orltom.dev/golang-http-example/internal/repository"
	"orltom.dev/golang-http-example/internal/resources"
	"orltom.dev/golang-http-example/internal/setup"
)

func main() {
	db, err := openDatabase()
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

func openDatabase() (*sql.DB, error) {
	envConfig, err := setup.LoadEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("could not load environment informations. %v", err)
	}

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", envConfig.DatabaseHost, envConfig.DatabasePort, envConfig.DatabaseUserName, envConfig.DatabasePassword, envConfig.DatabaseName)
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, fmt.Errorf("could not load database driver. %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database. %v", err)
	}
	return db, nil
}
