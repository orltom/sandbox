package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"orltom.dev/golang-http-example/internal/resources"
)

func Start(handler resources.JokeRestAPI) error {
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

	router.GET("/api/v1/jokes/random", handler.Random)
	router.GET("/api/v1/jokes/:UUID", handler.GetJokeByUUID)
	router.POST("/api/v1/jokes/", handler.Add)

	if err := router.Run(":8080"); err != nil {
		return fmt.Errorf("can not start web application. %v", err)
	}
	return nil
}
