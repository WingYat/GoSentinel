/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Repo: https://github.com/WingYat
 * @Date: 11/11/2023 - 1:56 am
 * @File: internal/controller/index.go
 * @Desc:
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexPage 渲染登录页面
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		// 如果有需要传递给模板的数据，可以在这里添加
	})
}
