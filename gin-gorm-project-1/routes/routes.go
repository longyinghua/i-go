package routes

import (
	"gin-gorm-app1/controller"
	"gin-gorm-app1/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//跨域中间件
	r.Use(middleware.CORSMiddleware())
	//用户注册
	r.POST("/api/auth/register", controller.Register)
	//用户登录
	r.POST("/api/auth/login", controller.Login)
	//获取用户信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	//解析token获取用户信息
	r.POST("/api/auth/information", controller.GetUserInfo)

	//删除用户
	r.POST("/api/user/delete", controller.DeleteUser)
	return r
}
