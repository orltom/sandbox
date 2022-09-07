package resources

import "github.com/gin-gonic/gin"

type JokeRestApi interface {
	GetJokeByUUID(c *gin.Context)
	Random(c *gin.Context)
	Add(c *gin.Context)
}
