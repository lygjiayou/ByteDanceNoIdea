package api

import (
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/service"
	"ByteDanceNoIdea/middleware"
	"ByteDanceNoIdea/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	//获取参数
	videoId, _ := strconv.ParseInt(c.Query("video_id"),10,64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	token := c.Query("token")
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

	//更新数据库
	if success := service.Action(userId, videoId, actionType); success == false {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_FAVORITEACTION_FALSE,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_FAVORITEACTION_FALSE),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: errmsg.LIKE_SUCCESS,
		StatusMsg: errmsg.GetErrMsg(errmsg.LIKE_SUCCESS),
	})	
	return 
}