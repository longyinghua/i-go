package main

import (
	"gin-gorm-app1/common"
	"gin-gorm-app1/dal/query"
	_ "gin-gorm-app1/docs"
	"gin-gorm-app1/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//// MySQLDSN MySQL data source name
//const MySQLDSN = "root:dfzy_12345@tcp(192.168.2.250:3306)/test?charset=utf8mb4&parseTime=True"
//
//func init() {
//	common.DB = common.ConnectDB(MySQLDSN).Debug() //  初始化数据库连接对象，返回数据库链接对象
//}

// 生成swagger注释
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name 龙应华
// @contact.url http://www.swagger.io/support
// @contact.email 542791872@qq.com
// @host 127.0.0.1:9090
// @BasePath /api/
func main() {
	InitConfig()

	//返回一个数据库连接对象
	common.DB = common.InitDB()

	//	设置默认DB对象
	query.SetDefault(common.DB)

	//初始化日志
	common.InitLogger()

	engine := gin.Default()
	//注册zap日志相关中间件
	engine.Use(common.GinLogger(), common.GinRecovery(true))

	//注册路由
	routes.CollectRoute(engine)

	//swagger路由控制
	var swagHandler gin.HandlerFunc
	swagHandler = ginSwagger.WrapHandler(swaggerfiles.Handler)

	if swagHandler != nil {
		engine.GET("/swagger/*any", swagHandler)
		//swagger访问地址 http://127.0.0.1:9090/swagger/index.html
	}

	//启动服务
	port := viper.GetString("server.port")
	if port != "" {
		panic(engine.Run(":" + port))
	}

	panic(engine.Run()) // listen and serve on 0.0.0.0:8080
}
