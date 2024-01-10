package routers

import (
	"github.com/gin-gonic/gin"
)
import "go-gorm-swagger-zap/middleware"

func LoadBookDeleteGroup1(engine *gin.Engine) {

	routerGroup2 := engine.Group("/api/v1/delete")

	{
		routerGroup2.POST("/books1", middleware.BookDelete1)

	}

}
