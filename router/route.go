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
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/index", service.GetIndex)

	r.GET("/user/getUserList", middleware.JWTHandler(), api.UserApi{}.GetUserList)
	r.POST("/user/createUser", middleware.JWTHandler(), api.UserApi{}.CreateUser)
	r.POST("/user/deleteUser", middleware.JWTHandler(), api.UserApi{}.DeleteUser)
	r.POST("/user/updateUser", middleware.JWTHandler(), api.UserApi{}.UpdateUser)
	r.POST("/user/login", api.UserApi{}.Login)

	//发送消息
	r.GET("/user/sendMsg", api.UserApi{}.SendMsg)
	return r
}
