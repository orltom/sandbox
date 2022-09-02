package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.GET("/joke", randomJoke)

	if err := router.Run(":8080"); err != nil {
		println("Can not start web application")
		os.Exit(1)
	}
}

func randomJoke(c *gin.Context) {
	msg := "hello world"

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
