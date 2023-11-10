/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Repo: https://github.com/WingYat
 * @Date: 10/11/2023 - 1:30 pm
 * @File: cmd/gosentinel/main.og.go
 * @Desc:
 */

package main

import (
	"GoSentinel/internal/config"
	"GoSentinel/internal/database"
	"GoSentinel/internal/router"
	"NetShieldControl/utils/log"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置、日志与数据库
	config.InitConfig()
	log.InitLogger()
	database.InitDatabase()

	r := router.InitRouter()
	r.Run(viper.GetString("server.http_address") + ":" + viper.GetString("server.http_port"))
}
