package server

import (
	"ByteDanceNoIdea/api"
	"ByteDanceNoIdea/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/public", "./public")

	// 无需token验证的路由放在该路由下
	publicApiRouter := r.Group("/douyin")
	{
		publicApiRouter.GET("/feed/", api.Feed)
		publicApiRouter.POST("/user/register/", api.UserRegister)
		publicApiRouter.POST("/user/login/", api.UserLogin)
		publicApiRouter.POST("/publish/action/", api.Publish)
		publicApiRouter.POST("/favorite/action/", api.FavoriteAction)
		//publicApiRouter.GET("/publish/list/", api.PublishList)
		//publicApiRouter.GET("/user/", api.UserInfo)
	}

	// 需要token验证的路由放在该路由下
	auth := r.Group("/douyin")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("/user/", api.UserInfo)
		auth.GET("/publish/list", api.GetPublishList)
		//auth.POST("/publish/action/", api.Publish),yanghai自己验证了，就不放在这里了
		auth.GET("/xxx/xxx/", api.TokenTest)
	}

	// basic apis
	//apiRouter.GET("/feed/", api.Feed)
	//apiRouter.GET("/user/", api.UserInfo)
	//apiRouter.POST("/user/register/", api.Register)
	//apiRouter.POST("/user/login/", api.Login)
	//apiRouter.POST("/publish/action/", api.Publish)
	//apiRouter.GET("/publish/list/", api.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", api.FavoriteAction)
	//apiRouter.GET("/favorite/list/", api.FavoriteList)
	//apiRouter.POST("/comment/action/", api.CommentAction)
	//apiRouter.GET("/comment/list/", api.CommentList)

	// extra apis - II
	//apiRouter.POST("/relation/action/", api.RelationAction)
	//apiRouter.GET("/relation/follow/list/", api.FollowList)
	//apiRouter.GET("/relation/follower/list/", api.FollowerList)
}
