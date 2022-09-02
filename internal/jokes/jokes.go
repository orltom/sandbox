package jokes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Random(c *gin.Context) {
	var msg = ""

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalln(err)
		msg = fmt.Sprintf("db error: %s", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		msg = fmt.Sprintf("db error: %s", err)
	}

	var joke string
	rows, err := db.Query("SELECT joke FROM chuck_norris ORDER BY RANDOM() LIMIT 1")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		msg = fmt.Sprintf("db error: %s", err)
	}

	for rows.Next() {
		err = rows.Scan(&joke)
		if err != nil {
			log.Fatalln(err)
			msg = fmt.Sprintf("db error: %s", err)
		}
		msg = joke
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

const (
	host     = "database.default"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "jokes"
)
