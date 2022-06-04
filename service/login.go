package service

import (
	"ByteDanceNoIdea/model"
)

// LoginService 登入服务
func LoginService(username string, password string) int {
	//var resp model.Response
	//user := model.User{
	//	UserName: req.UserName,
	//	Password: req.Password,
	//}
	//resp.StatusCode = int32(model.OK)
	//resp.StatusMsg = "user login success"
	// 登录验证
	code := model.CheckLogin(username, password)
	return code
}
