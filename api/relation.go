package api

import (
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/middleware"
	"ByteDanceNoIdea/utils/errmsg"
	"ByteDanceNoIdea/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	followTargetID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
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

	errno := service.FollowAction(actionType, followTargetID, userId)
	switch errno {
	case errmsg.ERROR_FOLLOW_FAIL:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_FOLLOW_FAIL,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_FOLLOW_FAIL),
		})
		return
	case errmsg.ERROR_CANCEL_FAIL:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_CANCEL_FAIL,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_FOLLOW_FAIL),
		})
		return
	case errmsg.ERROR_ALREADY_FOLLOW:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_ALREADY_FOLLOW,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_ALREADY_FOLLOW),
		})
		return
	case errmsg.ERROR_ALREADY_CANCEL:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR_ALREADY_CANCEL,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR_ALREADY_CANCEL),
		})
		return
	case errmsg.FOLLOW_SUCCESS:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.SUCCESS,
			StatusMsg: errmsg.GetErrMsg(errmsg.FOLLOW_SUCCESS),
		})
		return
	case errmsg.CANCEL_FOLLOW_SUCCESS:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.SUCCESS,
			StatusMsg: errmsg.GetErrMsg(errmsg.CANCEL_FOLLOW_SUCCESS),
		})
		return
	default:
		c.JSON(http.StatusOK, model.Response{
			StatusCode: errmsg.ERROR,
			StatusMsg: errmsg.GetErrMsg(errmsg.ERROR),
		})
		return
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
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

	userList, errno := service.FollowList(userId)
	if errno != false {
		c.JSON(http.StatusOK, model.RelationUserInfo{
			Response: model.Response{
					StatusCode:	errmsg.SUCCESS,
					StatusMsg:	errmsg.GetErrMsg(errmsg.SUCCESS),
			},
			UserInfos:	userList,
		})
	} else {
		c.JSON(http.StatusOK, model.RelationUserInfo{
			Response: model.Response{
					StatusCode:	errmsg.ERROR,
					StatusMsg:	errmsg.GetErrMsg(errmsg.ERROR),
			},
			UserInfos:	userList,
		})
	}
	
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
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

	userList, errno := service.FollowerList(userId)
	if errno != false {
		c.JSON(http.StatusOK, model.RelationUserInfo{
			Response: model.Response{
					StatusCode:	errmsg.SUCCESS,
					StatusMsg:	errmsg.GetErrMsg(errmsg.SUCCESS),
			},
			UserInfos:	userList,
		})
	} else {
		c.JSON(http.StatusOK, model.RelationUserInfo{
			Response: model.Response{
					StatusCode:	errmsg.ERROR,
					StatusMsg:	errmsg.GetErrMsg(errmsg.ERROR),
			},
			UserInfos:	userList,
		})
	}
}
