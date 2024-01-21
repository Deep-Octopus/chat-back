package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-chat/api"
	"go-chat/docs"
	"go-chat/middleware"
	"go-chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/index", service.GetIndex)

	r.GET("/user/getUserList", middleware.JWTHandler(), api.GetUserList)
	r.GET("/user/getUser", middleware.JWTHandler(), api.GetUser)
	r.GET("/user/getFriends", middleware.JWTHandler(), api.GetFriends)
	r.GET("/user/getGroups", middleware.JWTHandler(), api.GetGroups)
	r.GET("/user/getListMessage", middleware.JWTHandler(), api.GetListMessage)
	r.POST("/user/createUser", api.CreateUser)
	r.POST("/user/deleteUser", middleware.JWTHandler(), api.DeleteUser)
	r.POST("/user/updateUser", middleware.JWTHandler(), api.UpdateUser)
	r.POST("/user/login", api.Login)

	//发送消息
	r.GET("/message/sendMsg", api.SendMsg)
	r.GET("/message/sendUserMsg", api.SendUserMsg)

	r.POST("/message/getMessages", api.GetMessages)

	return r
}
