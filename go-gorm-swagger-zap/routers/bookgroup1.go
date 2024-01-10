package routers

import (
	"github.com/gin-gonic/gin"
)
import "go-gorm-swagger-zap/middleware"

func LoadBookGroup1(engine *gin.Engine) {

	routerGroup1 := engine.Group("/api/books/") //  路由分组 10.40.3.9:8080/api/books/v2    /api/books/v3

	//一个路由分组中可以有多个路由，通过一个{}包裹起来
	{
		routerGroup1.GET("/v1", middleware.BookInfo1())
		routerGroup1.GET("/v2", middleware.BookInfo2())
		routerGroup1.GET("/v3", middleware.BookInfo3())
		routerGroup1.GET("/v4", middleware.BookInfo4())
		routerGroup1.POST("/v5", middleware.BookInfo5)
		routerGroup1.POST("/v6", middleware.BookInfo6)
	}
}
