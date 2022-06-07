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


// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//获取参数
	var commentParam model.CommentRequestParam
	commentParam.VideoID, _ = strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentParam.ActionType, _ = strconv.Atoi(c.Query("action_type"))
	commentParam.CommentText = c.Query("comment_text")
	commentParam.CommentID, _ = strconv.ParseInt(c.Query("comment_id"), 10, 64)
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
	commentInfo, err := service.CommentActionService(commentParam, userId)
	switch err {
	case errmsg.ERROR_COMMENT_FAIL:
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_COMMENT_FAIL,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_COMMENT_FAIL),
		})
		return

	case errmsg.ERROR_DELETE_FAIL:
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_DELETE_FAIL,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_DELETE_FAIL),
		})
		return

	case errmsg.COMMENT_SUCCESS:
		c.JSON(http.StatusOK, model.CommentActionResponse{
			Response:	model.Response{
				StatusCode:errmsg.SUCCESS,
				StatusMsg:errmsg.GetErrMsg(errmsg.COMMENT_SUCCESS),
			},
			CommentReturn:	commentInfo,
		})
		return

	case errmsg.DELETE_COMMENT_SUCCESS:
		c.JSON(http.StatusOK, model.CommentActionResponse{
			Response:	model.Response{
				StatusCode:errmsg.SUCCESS,
				StatusMsg:errmsg.GetErrMsg(errmsg.DELETE_COMMENT_SUCCESS),
			},
		})
		return

	default:
		c.JSON(http.StatusOK,model.Response{
			StatusCode:errmsg.ERROR_COMMENT_FAIL,
			StatusMsg:errmsg.GetErrMsg(errmsg.ERROR_COMMENT_FAIL),
		})
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
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

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	c.JSON(http.StatusOK, model.CommentListResponse{
		Response: model.Response{
			StatusCode: errmsg.SUCCESS,
			StatusMsg: errmsg.GetErrMsg(errmsg.SUCCESS),
		},
		CommentList: service.CommentListService(userId, videoId),
	})
}
