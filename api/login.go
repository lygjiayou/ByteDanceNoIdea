package api

import (
	"douyin/model"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Login 登入api
func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: int32(model.ParamInvalid),
			StatusMsg:  "错误丫丫",
		})
	} else {
		username := c.Query("username")
		password := c.Query("password")
		req.Name = username
		req.Password = password
		resp := service.LoginService(&req)
		// 登入成功
		if resp.StatusCode == int32(model.OK) {
			//session := sessions.Default(c)
			//使用uuid作为token
			//UUID token=UUID.randomUUID;
			//token() := uuid.NewV4().String()

			// 根据Name从数据库查询ID
			//user := &model.User{
			//	ID:
			//}
			// 第一次登录需要根据用户名查询对应的ID，ID作为token的一部分,下次登录仍然重新生成token,token的作用是用户携带token访问其他资源时不用重新登录了
			user := model.User{
				Name: req.Name,
			}
			_, _ = user.FindByUsername()
			//id,exist := user.FindByUsername() // 根据用户名从数据库中查询用户ID
			//if exist != nil { // 这里不用判断了，上面service.LoginService已经判断过了
			//	c.JSON(http.StatusOK, model.Response{
			//		StatusCode: 1,
			//		StatusMsg:  "User doesn't exist",
			//	})
			//	return
			//}
			//u1 := uuid.NewV4().String()+req.Name+strconv.Itoa(int(id))
			u1 := uuid.NewV4().String() + req.Name
			fmt.Println(u1)
			//resp.Token = &u1
			//resp.UserID = &id
			c.JSON(http.StatusOK, model.LoginResponse{
				Response: model.Response{
					StatusCode: 0,
					StatusMsg:  "User login success",
				},
				Token:  u1,
				UserID: user.ID,
			})
			return
			//session.Clear()
			//session.Set("name", req.Name)
			//session.Save()
		}
		c.JSON(http.StatusBadRequest, resp)
	}
}

// 用户登出，因为客户端存放token,服务端不存放token,所以服务端不用操作，交给客户端来清除token即可

// 获取id

// Whoami 是get方法
func Whoami(c *gin.Context) {

}
