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

// Register 注册api
func Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: int32(model.ParamInvalid),
			StatusMsg:  "错误丫丫丫",
		})
	} else {
		username := c.Query("username")
		password := c.Query("password")
		req.Name = username
		req.Password = password
		if service.CheckRegisterParamService(&req) {
			// 参数校验通过，进行注册用户操作（注册操作前需要判断用户是否存在，满足条件则去注册用户）
			resp := service.RegisterService(&req)
			// 创建成功
			if resp.StatusCode == int32(model.OK) {
				// 第一次登录需要根据用户名查询对应的ID，ID作为token的一部分,下次登录仍然重新生成token,token的作用是用户携带token访问其他资源时不用重新登录了
				// 因为创建之后，ID在表中是自动+1的，所以创建之后需要根据用户名查询ID
				user := model.User{
					Name: req.Name,
				}
				_, err := user.FindByUsername() // 查询ID，这里本来作为token的一部分，后来不作为token的一部分了
				if err != nil {
					resp.StatusCode = int32(model.UnknownError)
					resp.StatusMsg = string("query failed")
				}
				//u1 := uuid.NewV4().String()+req.Name+strconv.Itoa(int(id))
				u1 := uuid.NewV4().String() + req.Name
				fmt.Println(u1)
				//resp.Token = &u1
				//resp.UserID = &id
				c.JSON(http.StatusOK, model.LoginResponse{
					Response: model.Response{
						StatusCode: 0,
						StatusMsg:  "User register success",
					},
					Token:  u1,
					UserID: user.ID,
				})
				return
			} else {
				// 校验不通过
				c.JSON(http.StatusOK, model.Response{
					StatusCode: int32(model.ParamInvalid),
					StatusMsg:  string("param doesn't correct"),
				})
			}
		}
	}
}

// 用户登出，因为客户端存放token,服务端不存放token,所以服务端不用操作，交给客户端来清除token即可
