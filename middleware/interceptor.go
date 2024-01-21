package middleware

import (
	"github.com/gin-gonic/gin"
	"go-chat/models"
	"go-chat/utils"
	"net/http"
	"strings"
)

//定义权限认证中间件
//jwt拦截器

func JWTHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var resp Response
		//引入jwt实现登录后的会话记录,登录会话发生登录完成之后
		//header获取token
		token := context.Request.Header.Get("Authorization")
		username := context.Request.Header.Get("Username")
		if token == "" {
			resp.Code = 302
			context.JSON(http.StatusOK, resp.WithMsg("请求未携带token无法访问!"))
			context.Abort()
		}
		cleanToken := strings.TrimPrefix(token, "Bearer ")

		//解析token
		claims, err := utils.ParseToken(cleanToken)
		if claims == nil || err != nil {
			resp.Code = 401
			context.JSON(http.StatusOK, resp.WithMsg("未携带有效token或已过期"))
			context.Abort()
		} else {
			user := models.TakeUserByUsername(username)
			if user == nil || user.ID != uint(claims.UserID) {
				//context.String(401, "未携带有效token或已过期")
				resp.Code = 401
				context.JSON(http.StatusOK, resp.WithMsg("未携带有效token或已过期"))
				context.Abort()
			}
			context.Next()

		}
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Username")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
