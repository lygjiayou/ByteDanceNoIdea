package api

import (
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/service"
	"ByteDanceNoIdea/middleware"
	"ByteDanceNoIdea/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTimeStr := c.Query("latest_time")
	var latestTime int64
	if latestTimeStr == "" {
		latestTime = time.Now().Unix()
	} else {
		latestTime, _ = strconv.ParseInt(c.Query("latest_time"), 10, 64)
	}

	var userId int64
	token := c.Query("token")
	if token == "" {
		userId = -1
	} else {
		//获取user id
		//解析username
		key, err := middleware.CheckToken(token)
		if err != errmsg.SUCCESS {
			userId = -1
		} else {
			//根据username获取userid
			user := model.User{
				UserName: key.Username,
			}
			var is_found error
			userId, is_found = user.FindByUsername()
			if is_found != nil {
				userId =  -1
			}
		}
	}

	videoInfos, nextTime := service.FeedService(latestTime, userId)
	c.JSON(http.StatusOK, model.FeedResponse{
		Response: model.Response{
			StatusCode:errmsg.SUCCESS,
			StatusMsg:errmsg.GetErrMsg(errmsg.SUCCESS),
		},
		VideoList: videoInfos,
		NextTime: nextTime,
	})
	return
}
