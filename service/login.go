package service

import (
	"douyin/model"
)

// LoginService 登入服务
func LoginService(req *model.LoginRequest) *model.Response {
	var resp model.Response
	user := model.User{
		Name:     req.Name,
		Password: req.Password,
	}
	rows, err := user.FindByUsernamePassword()
	if err != nil {
		resp.StatusCode = int32(model.UnknownError)
		return &resp
	}
	// 返回用户不存在错误
	if rows <= 0 {
		resp.StatusCode = int32(model.UnknownError)
		return &resp
	}
	// 返回密码错误
	if user.Password != req.Password {
		resp.StatusCode = int32(model.WrongPassword)
		return &resp
	}
	//resp.StatusCode = int32(model.OK)
	//resp.StatusMsg = "user login success"
	resp = model.Response(struct {
		StatusCode int32
		StatusMsg  string
	}{StatusCode: int32(model.OK), StatusMsg: string("user login success")})
	return &resp
}
