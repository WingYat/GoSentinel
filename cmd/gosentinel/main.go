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
	"flag"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// 定义命令行参数
	configPath := flag.String("config", "", "path to the config file")
	flag.Parse()

	// 检查配置文件路径是否已提供
	if *configPath == "" {
		log.Fatal("No config file path provided")
	}

	// 初始化配置
	config.InitConfig(*configPath)

	// 初始化日志与数据库
	//logger.InitLogger()
	database.InitDatabase()

	r := router.InitRouter()
	r.Run(viper.GetString("server.http_address") + ":" + viper.GetString("server.http_port"))

}
