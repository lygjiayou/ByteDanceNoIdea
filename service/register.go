package service

import (
	"douyin/model"
)

func CheckRegisterParamService(request *model.RegisterRequest) bool {
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
func RegisterService(req *model.RegisterRequest) *model.Response {
	var resp model.Response
	// 注册时，只提供用户名和密码即可，其他的默认为初始值比如int类型的默认为0,性别默认为male
	user := &model.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	// 先判断表中是否已经存在该用户
	rows, err := user.FindByName() // 根据名字查找用户是否存在
	if err != nil {
		resp.StatusCode = int32(model.UnknownError)
		return &resp
	}
	// 返回用户已注册错误
	if rows > 0 {
		resp.StatusCode = int32(model.UserHasExisted)
		return &resp
	}
	// 或者根据表中的unique特性，让数据库来帮助判断用户是否已经存在
	err2 := user.CreateUser()
	if err2 != nil {
		// 如果出现错误返回用户已经存在
		resp.StatusCode = int32(model.UserHasExisted)
		resp.StatusMsg = string("User has exist, create user failed!")
		return &resp
	}
	//else {
	//	resp.StatusCode = int32(model.OK)
	//	//resp.StatusMsg = string("the condition of create user meet")
	//}
	// 如果前面没有问题，就创建成功了
	resp = model.Response(struct {
		StatusCode int32
		StatusMsg  string
	}{StatusCode: int32(model.OK), StatusMsg: string("user register success")})
	return &resp
}
