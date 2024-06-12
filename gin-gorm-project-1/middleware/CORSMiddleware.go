package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域请求中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置允许跨域请求的域名
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许跨域请求的有效期
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置允许跨域请求的HTTP方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		// 设置允许跨域请求的请求头
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 设置允许跨域请求的凭证
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 如果请求方法为OPTIONS，则直接返回200状态码
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			// 否则继续执行后续操作
			ctx.Next()
		}
	}

}
