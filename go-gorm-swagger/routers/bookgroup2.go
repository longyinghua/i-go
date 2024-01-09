package routers

import (
	"github.com/gin-gonic/gin"
)
import "go-gorm/middleware"

func LoadBookGroup2(engine *gin.Engine) {

	routerGroup2 := engine.Group("/api/v1/")

	{
		routerGroup2.POST("/books1", middleware.BookInfo1())
		routerGroup2.POST("/books2", middleware.BookInfo2())
		routerGroup2.POST("/books3", middleware.BookInfo3())
		routerGroup2.POST("/books4", middleware.BookInfo4())
		routerGroup2.POST("/books5", middleware.BookInfo5)
		routerGroup2.POST("/books6", middleware.BookInfo6)
	}

}
