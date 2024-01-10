package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gorm-swagger-zap/dal"
	"go-gorm-swagger-zap/dal/query"
	_ "go-gorm-swagger-zap/docs"
	"go-gorm-swagger-zap/logger"
	"go-gorm-swagger-zap/routers"
	"log"
)

// MySQLDSN MySQL data source name
const MySQLDSN = "root:dfzy_12345@tcp(192.168.2.250:3306)/test?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnectDB(MySQLDSN).Debug() //  初始化数据库连接对象，返回数据库链接对象
}

//	@title			Book API
//	@version		1.0
//	@description	Book API+SQL 增删改查
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	龙应华
//	@contact.url	http://www.swagger.io/support
//	@contact.email	542791872@qq.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		10.40.3.9:8080
// @BasePath	/api/v1
func main() {
	//	设置默认DB对象
	query.SetDefault(dal.DB)

	logger.InitLogger()

	engine := gin.Default()
	//注册zap日志相关中间件
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	//路由组1
	//routers.LoadBookGroup1(engine)

	//路由组2
	//routers.LoadBookGroup2(engine)

	//路由组3
	routers.LoadBookCreateGroup1(engine)

	//路由组4，自定义SQL放置在统一的组中
	routers.LoadBookCustomGroup1(engine)

	//路由组5，update
	routers.LoadBookUpdateGroup1(engine)

	//路由组6，delete
	routers.LoadBookDeleteGroup1(engine)

	//swagger路由控制
	var swagHandler gin.HandlerFunc
	swagHandler = ginSwagger.WrapHandler(swaggerfiles.Handler)

	if swagHandler != nil {
		engine.GET("/swagger/*any", swagHandler)
	}

	err := engine.Run(":8080")

	if err != nil {
		log.Printf("start server error: %v", err)
		return
	}
}
