package router

import (
	"TikTok/controller"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func SetupRouter() {
	Router.GET("/douyin/feed", controller.Feed)
	Router.POST("/douyin/register", controller.Register)
	Router.POST("douyin/user/login", controller.Login)
	Router.GET("douyin/user", controller.User)
	Router.POST("douyin/publish/action", controller.Publish)
	Router.GET("douyin/publish/list", controller.PublishList)
}
