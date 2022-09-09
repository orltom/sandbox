package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"orltom.dev/golang-http-example/internal/resources"
	"orltom.dev/golang-http-example/internal/services"
)

func Start() error {
	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(gin.Logger())

	db, err := openDatabase()
	if err != nil {
		return fmt.Errorf("can not open connection to database. %v", err)
	}

	service := services.NewDatabaseJokeService(db)
	rest := resources.NewJokeRestResource(service)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.GET("/api/jokes/random", rest.Random)
	router.GET("/api/jokes/:UUID", rest.GetJokeByUUID)
	router.POST("/api/jokes/", rest.Add)

	if err := router.Run(":8080"); err != nil {
		log.Printf("Can not start web application. %v", err)
		return fmt.Errorf("can not start web application. %v", err)
	}
	return nil
}

func openDatabase() (*sql.DB, error) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
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
