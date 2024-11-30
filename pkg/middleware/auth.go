package middleware

import (
	"myshop/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		userID, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "无效的token"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
