/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Blog: https://www.gov-cn.cn
 * @Date: 9/11/2023 - 11:35 pm
 * @File: internal/config/config.go
 * @Desc:
 */

package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig(configFilePath string) {
	// 设置配置文件路径
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}
