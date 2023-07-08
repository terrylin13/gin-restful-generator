package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里进行身份验证和权限检查
		// 获取请求中的用户身份信息（例如令牌或会话）
		// 验证用户身份和权限
		// 如果用户没有足够的权限，可以返回错误响应或拒绝请求
		// 否则，将用户信息传递给后续的处理函数
		c.Set("user", &model.User{ID: 1, Name: "john", Roles: []string{"admin"}})
		c.Next()
	}
}
