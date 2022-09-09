package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewPostgresJokeRepository(db)
	uuid := uuid.New().String()
	rows := mock.NewRows([]string{"uuid", "joke"}).AddRow(uuid, "test")
	query := "SELECT .* FROM jokes WHERE uuid = \\$1"
	mock.ExpectQuery(query).WithArgs(uuid).WillReturnRows(rows)

	joke, err := repo.FindByID(uuid)

	assert.NotNil(t, joke)
	assert.NoError(t, err)
	assert.Equal(t, "test", joke.Joke)
	assert.Equal(t, uuid, joke.UUID)
}

func TestFindByIDError(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewPostgresJokeRepository(db)
	uuid := uuid.New().String()
	rows := mock.NewRows([]string{"uuid", "joke"}).AddRow(uuid, "test")
	query := "SELECT .* FROM jokes WHERE uuid = \\$1"
	mock.ExpectQuery(query).WithArgs(uuid).WillReturnRows(rows)

	joke, err := repo.FindByID("NOT_EXISTING_UUID")

	assert.Nil(t, joke)
	assert.Error(t, err)
}

func TestRandom(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewPostgresJokeRepository(db)
	uuid := uuid.New().String()
	rows := mock.NewRows([]string{"uuid", "joke"}).AddRow(uuid, "test")
	query := "SELECT .* FROM jokes ORDER BY RANDOM"
	mock.ExpectQuery(query).WillReturnRows(rows)

	joke, err := repo.Random()

	assert.NotNil(t, joke)
	assert.NoError(t, err)
	assert.Equal(t, "test", joke.Joke)
	assert.Equal(t, uuid, joke.UUID)
}

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("could not mock database. %v", err)
	}
	return db, mock
}
