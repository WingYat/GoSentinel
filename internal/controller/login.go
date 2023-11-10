/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Repo: https://github.com/WingYat
 * @Date: 10/11/2023 - 3:26 pm
 * @File: internal/controller/login.go
 * @Desc:
 */

package controller

import (
	"GoSentinel/internal/middleware"
	"GoSentinel/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
	// 导入其他需要的包
)

// LoginHandler 处理登录请求
func LoginHandler(c *gin.Context) {
	var form model.LoginForm

	// 绑定表单数据
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, model.JsonResponse{Code: -1, Data: nil, Msg: "Invalid request"})
		return
	}

	// 验证验证码
	captchaId := c.PostForm("captchaId")
	captchaValue := c.PostForm("captcha")
	if !VerifyCaptcha(captchaId, captchaValue) {
		fmt.Println("sb")
		c.JSON(http.StatusOK, model.JsonResponse{
			Code: -1,
			Data: nil,
			Msg:  "Captcha verification failed",
		})
		return
	}

	// TODO: 验证用户名和密码
	// 假设用户名和密码验证成功，并且获取到了userID
	userID := 123 // 假设的用户ID

	// 生成JWT令牌
	tokenString, err := middleware.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 设置JWT到Cookie
	expirationHours := viper.GetInt("jwt.expiration_hours")
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)
	middleware.SetTokenToCookie(c, tokenString, expirationTime)

	// 响应登录成功
	c.JSON(http.StatusOK, model.JsonResponse{
		Code: 0,
		Data: gin.H{"message": "Login successful"},
		Msg:  "Login successful",
	})
}

// IndexPage 渲染登录页面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		// 如果有需要传递给模板的数据，可以在这里添加
	})
}
