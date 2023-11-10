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
	"strconv"
	"time"
	// 其他导入...
)

// JwtMiddleware 创建JWT认证中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := getTokenFromCookie(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// 验证token逻辑...
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 需要返回用于验证JWT的密钥
			jwtSecret := viper.GetString("jwt.secret")
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

// GenerateJWT 生成JWT令牌
func GenerateJWT(userID int) (string, error) {
	// 从配置文件中读取密钥和令牌有效期
	jwtSecret := viper.GetString("jwt.secret")
	expirationHours := viper.GetInt("jwt.expiration_hours")
	if expirationHours <= 0 {
		expirationHours = 24 // 默认为24小时
	}

	// 设置令牌有效期
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	// 创建JWT声明
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(userID),
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并生成最终的JWT令牌
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// getTokenFromCookie 从Cookie中提取JWT令牌
func getTokenFromCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return "", err
	}
	return cookie, nil
}

// SetTokenToCookie 设置JWT到Cookie
func SetTokenToCookie(c *gin.Context, tokenString string, expiration time.Time) {
	c.SetCookie("jwt", tokenString, int(expiration.Sub(time.Now()).Seconds()), "/", "", false, true)
}

func CheckJWTAndRespond() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := getTokenFromCookie(c)
		if err != nil {
			// 如果没有找到cookie或cookie无效，重定向到登录页面
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort() // 中断请求处理链
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			jwtSecret := viper.GetString("jwt.secret")
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			// 如果JWT解析失败或无效，重定向到登录页面
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort() // 中断请求处理链
			return
		}

		// JWT有效，设置用户信息并继续
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["id"])
		c.Next()
	}
}
