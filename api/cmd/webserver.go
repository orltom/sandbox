package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"orltom.dev/golang-http-example/internal/resources"
)

func Start(handler resources.JokeRestAPI) error {
	gin.DisableConsoleColor()
	gin.EnableJsonDecoderDisallowUnknownFields()

	appRouter := gin.Default()
	appRouter.Use(gin.Logger())

	metricRouter := gin.Default()
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.UseWithoutExposingEndpoint(appRouter)
	m.Expose(metricRouter)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	appRouter.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	appRouter.GET("/api/v1/jokes/random", handler.Random)
	appRouter.GET("/api/v1/jokes/:UUID", handler.GetJokeByUUID)
	appRouter.POST("/api/v1/jokes/", handler.Add)

	// Run metric endpoint on different port.
	go func() {
		err := metricRouter.Run(":9090")
		if err != nil {
			log.Printf("can not start metric. %v", err)
		}
	}()

	if err := appRouter.Run(":8080"); err != nil {
		return fmt.Errorf("can not start web application. %v", err)
	}
	return nil
}
