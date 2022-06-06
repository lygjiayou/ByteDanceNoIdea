package api

import (
	"ByteDanceNoIdea/middleware"
	"ByteDanceNoIdea/model"
	"ByteDanceNoIdea/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

func UserRegister(c *gin.Context) {
	var data model.User
	var token string
	data.UserName = c.Query("username")
	data.Password = c.Query("password")

	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		// 数据库中创建用户
		data.ID, code = model.CreateUser(&data)

		// 生成token
		token, code = middleware.SetToken(data.UserName, data.Password)
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": code,
		"user_id":     data.ID,
		"status_msg":  errmsg.GetErrMsg(code),
		"token":       token,
	})
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func UserLogin(c *gin.Context) {
	var data model.User
	var token string
	data.UserName = c.Query("username")
	data.Password = c.Query("password")
	//data.Password = model.ScryptPW(data.Password)
	// 根据username查询userid
	data.ID, _ = data.FindByUsername()
	code = model.CheckLogin(data.UserName, data.Password)
	if code == errmsg.SUCCESS {
		// 生成token
		token, code = middleware.SetToken(data.UserName, data.Password)
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: errmsg.GetErrMsg(code)},
		UserId:   data.ID,
		Token:    token,
	})

}

//type UserInfoResponse struct {
//	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
//	StatusMsg  string `json:"status_msg"`  // 返回状态描述
//	User       model.User   `json:"user"`        // 用户信息
//}

func UserInfo(c *gin.Context) {
	//var resp UserInfoResponse
	var user model.User
	//id := c.Query("user_id")
	s := c.Query("user_id")
	strconv.Atoi(s)
	_, _ = user.FindByUserID(s)
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{StatusCode: 0, StatusMsg: "success"},
		User:     user,
	})
	//var user model.User
	//user.UserName = key.Username
	//// 通过token解析出的username检查user的合法性
	//if len(user.UserName) < 6 || len(user.UserName) > 32 {
	//	c.JSON(http.StatusOK, api.Response{
	//		StatusCode: 1,
	//		StatusMsg:  "User doesn't exist",
	//	})
	//}
	//// 通过token解析出的username判断username是否还存在
	//rows := user.FindByName()
	//if rows == 0 {
	//	c.JSON(http.StatusOK, model.Response{
	//		StatusCode: 1,
	//		StatusMsg:  "User doesn't exist",
	//	})
	//}
	//_, err := user.FindByTokename() // 根据user.UserName查询user所有的信息
	//if err != nil {
	//	resp.StatusCode = int64(model.UnknownError)
	//	resp.StatusMsg = "query failed"
	//}
	//c.JSON(http.StatusOK, model.UserInfoResponse{
	//	StatusCode: 0,
	//	User:       user,
	//})
}

func TokenTest(c *gin.Context) {
	fmt.Println("username: ")
	fmt.Println(c.Keys["username"])
	c.JSON(http.StatusOK, gin.H{
		"username": c.Keys["username"],
	})
}

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"sync/atomic"
//)
//
//// usersLoginInfo use map to store user info, and key is username+password for demo
//// user data will be cleared every time the server starts
//// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}
