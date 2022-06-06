package api

import (
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPublishList(c *gin.Context) {
	var resp model.PublishListResponse
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(http.StatusOK, resp)
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = "success"
		resp.VideoList = service.PublishListService(model.PublishListRequest{UserID: int64(userID)})
		c.JSON(http.StatusOK, resp)
	}

}
