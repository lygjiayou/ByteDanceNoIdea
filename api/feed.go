package api

import (
	"ByteDanceNoIdea/service"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	service.FeedService(c)
}
