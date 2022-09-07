package resources

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"orltom.dev/golang-http-example/internal/services"
)

type jokeRestResource struct {
	service services.JokeService
}

var _ JokeRestApi = &jokeRestResource{}

func NewJokeRestResource(service services.JokeService) jokeRestResource {
	return jokeRestResource{service: service}
}

func (r *jokeRestResource) GetJokeByUUID(c *gin.Context) {
	var uuid = c.Param("UUID")
	joke, err := r.service.Get(uuid)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not find joke",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, joke)
}

func (r *jokeRestResource) Random(c *gin.Context) {
	joke, err := r.service.Random()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not find joke",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, joke)
}

func (r *jokeRestResource) Add(c *gin.Context) {
	c.JSON(http.StatusCreated, "OK")
}
