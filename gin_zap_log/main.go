package main

import (
	"fmt"
	"gin_zap_log/config"
	"gin_zap_log/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	//load config from config.json
	if len(os.Args) < 1 {
		return
	}

	logConfig := config.LogConfig{
		Level:      "info",
		Filename:   "D:\\code\\golang\\code\\gin_zap_log\\ginzap.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 3,
	}

	//init logger
	err := logger.InitLogger(logConfig)
	if err != nil {
		fmt.Printf("init logger failed, err:%v", err)
		return
	}

	//gin.SetMode(config.Conf.Mode)

	engine := gin.Default()
	//注册zap相关中间件
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	engine.GET("/hello", func(context *gin.Context) {
		var (
			name = "long"
			age  = 30
		)

		//记录日志并使用zap.xxx(key,val)记录相关字段
		zap.L().Debug(
			"hello world",
			zap.String("user", name),
			zap.Int("age", age),
		)

		context.JSON(
			http.StatusOK,
			gin.H{
				"hello": "world",
			},
		)
	})

	engine.Run(":8080")
}
