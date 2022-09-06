package jokes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJokeByUUID(c *gin.Context) {
	var uuid = c.Param("UUID")

	db, err := ConnectDatabase()
	defer db.Close()

	if err != nil {
		dbErrorHandling(c, err)
		return
	}

	rows, err := db.Query("SELECT joke FROM jokes WHERE uuid = $1 LIMIT 1", uuid)
	defer rows.Close()
	if err != nil {
		dbErrorHandling(c, err)
		return
	}

	var joke string
	for rows.Next() {
		err = rows.Scan(&joke)
		if err != nil {
			dbErrorHandling(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid":    uuid,
		"message": joke,
	})
}

func Random(c *gin.Context) {
	db, err := ConnectDatabase()
	defer db.Close()

	if err != nil {
		dbErrorHandling(c, err)
		return
	}

	rows, err := db.Query("SELECT uuid, joke FROM jokes ORDER BY RANDOM() LIMIT 1")
	defer rows.Close()
	if err != nil {
		dbErrorHandling(c, err)
		return
	}

	var joke string
	var uuid string
	for rows.Next() {
		err = rows.Scan(&uuid, &joke)
		if err != nil {
			dbErrorHandling(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid":    uuid,
		"message": joke,
	})
}

func dbErrorHandling(c *gin.Context, err error) {
	log.Fatalf("DB Issue: %e", err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal DB Issue",
	})
}

func ConnectDatabase() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

const (
	host     = "database.default"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "jokes"
)
