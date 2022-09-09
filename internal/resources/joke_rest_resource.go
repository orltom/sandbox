package resources

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"orltom.dev/golang-http-example/internal/services"
)

type jokeRestResource struct {
	service services.JokeService
}

var _ JokeRestAPI = &jokeRestResource{}

func NewJokeRestResource(service services.JokeService) JokeRestAPI {
	return &jokeRestResource{service: service}
}

func (r *jokeRestResource) GetJokeByUUID(c *gin.Context) {
	var uuid = c.Param("UUID")
	joke, err := r.service.Get(uuid)
	if err != nil {
		log.Printf("Could not get joke from database. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find joke",
		})
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, joke)
}

func (r *jokeRestResource) Random(c *gin.Context) {
	joke, err := r.service.Random()
	if err != nil {
		log.Printf("Could not get random joke from database. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find joke",
		})
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, joke)
}

func (r *jokeRestResource) Add(c *gin.Context) {
	joke, err := r.service.Add("bla")
	if err != nil {
		log.Printf("Could not store joke to database. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not add joke",
		})
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusCreated, joke)
}
