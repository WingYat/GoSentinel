/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Repo: https://github.com/WingYat
 * @Date: 10/11/2023 - 8:12 pm
 * @File: internal/controller/captcha.go
 * @Desc:
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成并返回验证码
func GenerateCaptcha(c *gin.Context) {
	driver := base64Captcha.DefaultDriverDigit
	captcha := base64Captcha.NewCaptcha(driver, store)
	captchaId, captchaImage, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成验证码失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"captchaId": captchaId, "captchaImage": captchaImage})
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(captchaId, captchaValue string) bool {
	return store.Verify(captchaId, captchaValue, true)
}
