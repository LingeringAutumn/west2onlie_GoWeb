package middlewares

import (
	"MyTodoList/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未提供令牌"})
			c.Abort()
			return
		}
		userID, err := config.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "无效令牌"})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
