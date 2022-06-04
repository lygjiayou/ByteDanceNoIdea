package api

import (
	"douyin/model"
	"github.com/gin-gonic/gin"
)

func GetPublishList(c *gin.Context) {
	var resp model.PublishListResponse
	userID := c.Query("user_id")
	var videoList []model.Video

}
