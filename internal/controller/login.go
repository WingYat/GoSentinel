/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Repo: https://github.com/WingYat
 * @Date: 10/11/2023 - 3:26 pm
 * @File: internal/controller/login.go
 * @Desc:
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginPage 渲染登录页面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		// 如果有需要传递给模板的数据，可以在这里添加
	})
}
