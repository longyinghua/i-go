package routers

import (
	"github.com/gin-gonic/gin"
)
import "go-gorm-swagger-zap/middleware"

func LoadBookUpdateGroup1(engine *gin.Engine) {

	routerGroup2 := engine.Group("/api/v1/update")

	{
		routerGroup2.POST("/books1", middleware.BookUpdate1)
		routerGroup2.POST("/books2", middleware.BookUpdate2)
		routerGroup2.POST("/books3", middleware.BookUpdate3)
		routerGroup2.POST("/books4", middleware.BookUpdate4)
	}

}
