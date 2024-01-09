package main

import (
	"github.com/gin-gonic/gin"
	"go-gorm/dal"
	"go-gorm/dal/query"
	"go-gorm/routers"
	"log"
)

// MySQLDSN MySQL data source name
const MySQLDSN = "root:dfzy_12345@tcp(192.168.2.250:3306)/test?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnectDB(MySQLDSN).Debug() //  初始化数据库连接对象，返回数据库链接对象
}

func main() {
	//	设置默认DB对象
	query.SetDefault(dal.DB)

	engine := gin.Default()

	//路由组1
	//routers.LoadBookGroup1(engine)

	//路由组2
	routers.LoadBookGroup2(engine)

	//路由组3
	routers.LoadBookCreateGroup1(engine)

	//路由组4，自定义SQL放置在统一的组中
	routers.LoadBookCustomGroup1(engine)

	//路由组5，update
	routers.LoadBookUpdateGroup1(engine)

	//路由组6，delete
	routers.LoadBookDeleteGroup1(engine)

	err := engine.Run(":8080")

	if err != nil {
		log.Printf("start server error: %v", err)
		return
	}
}
