package routers

import (
	"github.com/gin-gonic/gin"
)
import "go-gorm-swagger-zap/middleware"

func LoadBookCustomGroup1(engine *gin.Engine) {

	routerGroup1 := engine.Group("/api/v1/custom/")

	{
		routerGroup1.POST("/books1", middleware.CustomSelect1())
		routerGroup1.POST("/books2", middleware.CustomSelect2)
		routerGroup1.POST("/books3", middleware.CustomSelect3)
		routerGroup1.POST("/books4", middleware.CustomSelect4)
		routerGroup1.POST("/books5", middleware.CustomSelect5)
		routerGroup1.POST("/books6", middleware.CustomSelect6)
	}

}
