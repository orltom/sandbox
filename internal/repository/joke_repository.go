package repository

import (
	"database/sql"
	"fmt"
)

type postgresJokeRepository struct {
	db *sql.DB
}

var _ JokeRepository = &postgresJokeRepository{}

func NewPostgresJokeRepository(db *sql.DB) JokeRepository {
	return &postgresJokeRepository{db: db}
}

func (s *postgresJokeRepository) FindByID(uuid string) (*Joke, error) {
	query := "SELECT uuid, joke FROM jokes WHERE uuid = $1 LIMIT 1"
	row := s.db.QueryRow(query, uuid)

	var joke Joke
	err := row.Scan(&joke.UUID, &joke.Joke)
	if err != nil {
		return nil, fmt.Errorf("could not find joke. Reason: %v", err)
	}
	return &joke, nil
}

func (s *postgresJokeRepository) Create(uuid string, joke string) (*Joke, error) {
	query := "INSERT INTO jokes (uuid, joke) VALUES ($1, $2)"
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("could not create database transaction. Reason: %v", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("invalid SQL statement. Reason: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid, joke)
	if err != nil {
		return nil, fmt.Errorf("could not add joke to database. Reason: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("could not add joke to database. Reason: %v", err)
	}

	stmt.QueryRow(uuid, joke)
	return s.FindByID(uuid)
}

func (s *postgresJokeRepository) Random() (*Joke, error) {
	query := "SELECT uuid, joke FROM jokes ORDER BY RANDOM() LIMIT 1"
	row := s.db.QueryRow(query)

	var joke Joke
	err := row.Scan(&joke.UUID, &joke.Joke)
	if err != nil {
		return nil, fmt.Errorf("could not map database to joke. Reason: %v", err)
	}
	return &joke, nil
}
