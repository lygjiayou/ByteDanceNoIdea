package service

import (
	"ByteDanceNoIdea/model"
)

func CheckRegisterParamService(request *model.User) bool {
	username := request.UserName
	password := request.Password
	// 用户名长度不能大于32位字符
	if len(username) < 6 || len(username) > 32 {
		return false
	}
	// 密码大于5位，不超过32位
	if len(password) < 6 || len(password) > 32 {
		return false
	}
	return true
}

// RegisterService 注册服务
//func RegisterService(user *model.User) int {
//	//var resp model.Response
//	// 注册时，只提供用户名和密码即可，其他的默认为初始值比如int类型的默认为0,性别默认为male
//	// 先判断表中是否已经存在该用户
//	code := user.FindByName() // 根据名字查找用户是否存在
//	return code
//}
