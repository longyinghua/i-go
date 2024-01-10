package routers

import (
	"github.com/gin-gonic/gin"
)
import "go-gorm-swagger-zap/middleware"

func LoadBookCreateGroup1(engine *gin.Engine) {

	routerGroup1 := engine.Group("/api/v1/create")

	{
		//单个插入
		routerGroup1.POST("/books1", middleware.BookCreate1())
		routerGroup1.POST("/books2", middleware.BookCreate2)
		//	批量插入,两种模型绑定方式
		routerGroup1.POST("/books3", middleware.BookCreate3)
		routerGroup1.POST("/books4", middleware.BookCreate4)
	}

}
