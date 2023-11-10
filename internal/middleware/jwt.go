/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Blog: https://www.gov-cn.cn
 * @Date: 10/11/2023 - 12:17 am
 * @File: internal/middleware/jwt.go
 * @Desc:
 */

package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	// 其他导入...
)

// JwtMiddleware 创建JWT认证中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// 验证token逻辑...
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 需要返回用于验证JWT的密钥
			jwtSecret := viper.GetString("jwt_secret")
			return []byte(jwtSecret), nil
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

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取JWT令牌
		jwtSecret := viper.GetString("jwt_secret")
		tokenString := getTokenFromHeader(c)

		// 检查令牌是否存在
		if tokenString == "" {
			redirectToLogin(c)
			return
		}

		// 验证JWT令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 返回用于验证的密钥
			// 注意: 确保根据您的实际情况修改密钥获取方式
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			redirectToLogin(c)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

// getTokenFromHeader 从请求头中提取JWT令牌
func getTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return ""
	}
	return parts[1]
}

// redirectToLogin 重定向到登录页面
func redirectToLogin(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/login")
	c.Abort()
}
