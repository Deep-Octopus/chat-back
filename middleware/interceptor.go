package middleware

import (
	"github.com/gin-gonic/gin"
	"go-chat/utils"
)

//定义权限认证中间件
//jwt拦截器

func JWTHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

		//引入jwt实现登录后的会话记录,登录会话发生登录完成之后
		//header获取token
		token := context.Request.Header.Get("token")
		if token == "" {
			context.String(302, "请求未携带token无法访问!")
			context.Abort()
		}
		//解析token
		claims, err := utils.ParseToken(token)
		if claims == nil || err != nil {
			context.String(401, "未携带有效token或已过期")
			context.Abort()
		} else {
			context.Next()

		}
	}
}
