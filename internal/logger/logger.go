/**
 * @Auth: WingYat - win9yat.cheung@gmail.com
 * @Blog: https://www.gov-cn.cn
 * @Date: 9/11/2023 - 11:38 pm
 * @File: internal/logger/logger.go
 * @Desc:
 */

package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

// LogFormatter 自定义日志格式
type LogFormatter struct{}

// Format 格式化日志输出
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level.String(), entry.Message)

	return []byte(logMessage), nil
}

func InitLogger() {
	logLevel, err := logrus.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	// 设置日志格式
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logLevel)

	// 设置日志输出（普通日志）
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/app-server.log",
		MaxSize:    10, // MB
		MaxBackups: 3,
		MaxAge:     viper.GetInt("logging.max_age_days"),
		Compress:   true,
	})

	// 错误日志输出
	errorLogger := logrus.New()
	errorLogger.SetFormatter(&LogFormatter{})
	errorLogger.SetLevel(logrus.ErrorLevel)
	errorLogger.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/app-error.log",
		MaxSize:    10, // MB
		MaxBackups: 3,
		MaxAge:     viper.GetInt("logging.max_age_days"),
		Compress:   true,
	})

	// 将 errorLogger 设置为全局的错误日志处理器
	// 示例：errorLogger.Error("这是一条错误日志")
}
