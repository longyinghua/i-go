package main

import (
	"gin-gorm-app1/common"
	"gin-gorm-app1/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	InitConfig()

	//返回一个数据库连接对象
	common.InitDB()

	//初始化日志
	common.InitLogger()

	engine := gin.Default()
	//注册zap日志相关中间件
	engine.Use(common.GinLogger(), common.GinRecovery(true))

	//注册路由
	routes.CollectRoute(engine)

	//启动服务
	port := viper.GetString("server.port")
	if port != "" {
		panic(engine.Run(":" + port))
	}

	panic(engine.Run()) // listen and serve on 0.0.0.0:8080
}
