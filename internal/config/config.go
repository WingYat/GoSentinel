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

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}
