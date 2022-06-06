package api

import (
	"douyin/model"
	"douyin/service"
	"douyin/middleware"
	"douyin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	//获取token，鉴权并解析userid
	token := c.PostForm("token")
	//token不存在
	if token == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_TOKEN_EXIST,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_TOKEN_EXIST),
		})
		return
	}

	//获取user id
	//解析username
	key,err := middleware.CheckToken(token)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, model.Response{
			StatusCode:errmsg.ERROR_TOKEN_WRONG,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_TOKEN_WRONG),
		})
		return
	}

	//根据username获取userid
	user := model.User{
		UserName:key.Username,
	}
	userId, is_found :=  user.FindByUsername()
	if(is_found != nil) {
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_USER_NOT_EXIST,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_EXIST),
		})
		return
	}

	//获取操作视频id和操作
	videoId, _ := strconv.ParseInt(c.PostForm("video_id"),10,64)
	actionType, _ := strconv.Atoi(c.PostForm("action_type"))

	if success := service.Action(userId, videoId, actionType); success == false {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_FAVORITEACTION_FALSE,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_FAVORITEACTION_FALSE),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: errmsg.SUCCESS,
		StatusMsg: errmsg.GetErrMsg(errmsg.SUCCESS),
	})	
	return 
}