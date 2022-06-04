package api

import (
	"douyin/middleware"
	"douyin/model"
	"douyin/utils/errmsg"

	//"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	//var req model.UserRequest
	var resp model.UserInfoResponse
	token := c.Query("token")
	// 验证token
	key, code := middleware.CheckToken(token)
	if code == errmsg.ERROR {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "token is failed",
		})
	}
	// 解析token  -- 这是错的
	// 因为生成的token：uuid+name,而uuid的长度是36位固定的，
	// 取token的36位开始的字符串后面就是name,然后根据name从数据库中查询该用户的信息，作为user
	//s := token[36:]
	//req.Token = token
	var user model.User
	user.UserName = key.Username
	// 通过token解析出的username检查user的合法性
	if len(user.UserName) < 6 || len(user.UserName) > 32 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}
	// 通过token解析出的username判断username是否还存在
	rows, _ := user.FindByName()
	if rows == 0 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}
	_, err := user.FindByTokename() // 根据user.UserName查询user所有的信息
	if err != nil {
		resp.StatusCode = int64(model.UnknownError)
		resp.StatusMsg = "query failed"
	}
	c.JSON(http.StatusOK, model.UserInfoResponse{
		StatusCode: 0,
		User:       user,
	})
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
}
