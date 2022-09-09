package resources

import "github.com/gin-gonic/gin"

type JokeRestAPI interface {
	GetJokeByUUID(c *gin.Context)
	Random(c *gin.Context)
	Add(c *gin.Context)
}

type SaveJoke struct {
	Joke string `json:"joke,omitempty"`
}
