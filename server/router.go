package server

import (
	"douyin/api"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	// 无需token验证的路由放在该路由下
	publicApiRouter := r.Group("/douyin")
	{
		publicApiRouter.POST("/user/register/", api.Register)
		publicApiRouter.POST("/user/login/", api.Login)
		publicApiRouter.GET("/feed/", controller.Feed)

		publicApiRouter.POST("/publish/action/", controller.Publish)
		publicApiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		publicApiRouter.POST("/favorite/action/", controller.FavoriteAction)
		publicApiRouter.GET("/favorite/list/", controller.FavoriteList)
		publicApiRouter.POST("/comment/action/", controller.CommentAction)
		publicApiRouter.GET("/comment/list/", controller.CommentList)

		// extra apis - II
		publicApiRouter.POST("/relation/action/", controller.RelationAction)
		publicApiRouter.GET("/relation/follow/list/", controller.FollowList)
		publicApiRouter.GET("/relation/follower/list/", controller.FollowerList)

	}
	// 需要token验证的路由放在该路由下
	auth := r.Group("/douyin")
	//auth.Use(middleware.JwtToken())
	{
		auth.GET("/user/", api.UserInfo)
		auth.GET("/publish/list", api.GetPublishList)
	}
	// basic apis

}
