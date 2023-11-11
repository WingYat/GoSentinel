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
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
		c.JSON(http.StatusUnauthorized, model.JsonResponse{Code: -1, Data: nil, Msg: "Captcha verification failed"})
		return
	}
	// 获取用户信息
	user, err := model.GetUserByUsername(form.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户名不存在
			c.JSON(http.StatusUnauthorized, model.JsonResponse{-1, nil, "Username does not exist"})
		} else {
			// 其他数据库错误
			c.JSON(http.StatusInternalServerError, model.JsonResponse{-1, nil, "Internal server error"})
		}
		return
	}

	// 检查密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password)); err != nil {
		// 密码不匹配
		c.JSON(http.StatusUnauthorized, model.JsonResponse{-1, nil, "Invalid password"})
		return
	}

	// 生成JWT令牌，使用数据库中查询到的用户ID
	tokenString, err := middleware.GenerateJWT(user.ID) // 使用user.ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.JsonResponse{-1, nil, "Failed to generate token"})
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
