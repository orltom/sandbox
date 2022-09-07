package services

import (
	"database/sql"
	"fmt"
)

type databaseJokeService struct {
	db *sql.DB
}

var _ JokeService = &databaseJokeService{}

func NewDatabaseJokeService(db *sql.DB) *databaseJokeService {
	return &databaseJokeService{db: db}
}

func (s *databaseJokeService) Get(uuid string) (*Joke, error) {
	query := "SELECT uuid, joke FROM jokes WHERE uuid = $1 LIMIT 1"
	row := s.db.QueryRow(query, uuid)

	var joke Joke
	err := row.Scan(&joke.UUID, &joke.Joke)
	if err != nil {
		return nil, fmt.Errorf("could not map database to joke. Reason: %e", err)
	}
	return &joke, nil
}

func (s *databaseJokeService) Add(_ string) (*Joke, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (s *databaseJokeService) Random() (*Joke, error) {
	query := "SELECT uuid, joke FROM jokes ORDER BY RANDOM() LIMIT 1"
	row := s.db.QueryRow(query)

	var joke Joke
	err := row.Scan(&joke.UUID, &joke.Joke)
	if err != nil {
		return nil, fmt.Errorf("could not map database to joke. Reason: %e", err)
	}
	return &joke, nil
}
