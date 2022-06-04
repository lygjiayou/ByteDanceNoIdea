package api

import (
	"douyin/model"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPublishList(c *gin.Context) {
	var resp model.PublishListResponse
	//userID := c.Query("user_id")
	var user model.User
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(http.StatusOK, resp)
	} else {
		user.ID = int64(userID)
		//根据userID获取author信息并填入user
		user.FindUserInfoByID()
		resp.StatusCode = 0
		resp.StatusMsg = "OK"
		resp.VideoList = service.PublishListService(model.PublishListRequest{UserID: user.ID})
		c.JSON(http.StatusOK, resp)
	}

}
