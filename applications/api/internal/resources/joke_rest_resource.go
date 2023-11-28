package resources

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"orltom.dev/golang-http-example/internal/repository"
)

type jokeRestResource struct {
	repository repository.JokeRepository
}

var _ JokeRestAPI = &jokeRestResource{}

func NewJokeRestResource(service repository.JokeRepository) JokeRestAPI {
	return &jokeRestResource{repository: service}
}

func (r *jokeRestResource) GetJokeByUUID(c *gin.Context) {
	var uuid = c.Param("UUID")
	joke, err := r.repository.FindByID(uuid)
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
	joke, err := r.repository.Random()
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
	newJoke := SaveJoke{}
	if err := c.BindJSON(&newJoke); err != nil {
		log.Printf("Invalid JSON request. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON input",
		})
		c.Abort()
		return
	}

	uuid := uuid.New().String()
	joke, err := r.repository.Create(uuid, newJoke.Joke)
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
