/**
 * @Auth: yycheung - win9yat.cheung@gmail.com
 * @Blog: https://www.gov-cn.cn
 * @Date: 10/11/2023 - 12:17 am
 * @File: internal/middleware/jwt.go
 * @Desc:
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// 其他导入...
)

// JwtMiddleware 创建JWT认证中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// 验证token逻辑...
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 返回密钥或公钥
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// 设置用户信息到上下文
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["id"])

		c.Next()
	}
}
