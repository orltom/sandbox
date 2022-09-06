package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"orltom.dev/golang-http-example/internal/jokes"
)

func Start() {
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

	router.GET("/api/jokes/random", jokes.Random)
	router.GET("/api/jokes/:UUID", jokes.GetJokeByUUID)

	if err := router.Run(":8080"); err != nil {
		println("Can not start web application")
		os.Exit(1)
	}
}
