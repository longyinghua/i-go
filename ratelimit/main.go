package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// 创建一个流控器，使用时间粒度为1秒，初始容量为100，最大容量为100的速率限制器
var limiter = ratelimit.NewBucketWithQuantum(time.Second, 100, 100)

/*
有下面特点:
令牌桶初始化后里面就有 100 个令牌
每秒钟会产生 100 个令牌, 保证每秒最多有 100 个请求通过限流器, 也就是说 QPS 的上限是 100
流量过大时能够启动限流, 在限流过程中, 仍然能让部分流量通过
go-stress-testing -c 10 -n 10 -u http://10.40.3.9:8080/hello 通过压测可看到结果，每秒大于100个请求就会出现429错误
*/

// tokenRateLimiter 生成一个中间件函数，用于实现令牌桶限流算法。
//
//	在新的处理函数中，首先尝试从令牌桶中获取一个可用的令牌(limiter.TakeAvailable(1))。如果获取到的令牌数量为0,说明请求过于频繁，此时返回HTTP状态码429和错误信息"Too Many Requests"。否则，将剩余可用令牌数量和令牌桶的总容量分别设置到响应头的"X-RateLimit-Remaining"和"X-RateLimit-Limit"字段中，然后调用context.Next()继续处理后续的路由
func tokenRateLimiter() gin.HandlerFunc {
	fmt.Println("token create rate:", limiter.Rate())     //  打印当前的令牌创建速率(limiter.Rate())
	fmt.Println("available token :", limiter.Available()) //  和可用令牌数量(limiter.Available())
	return func(context *gin.Context) {
		if limiter.TakeAvailable(1) == 0 {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": 429,
				"msg":  "Too Many Requests",
			})
		} else {
			context.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
			context.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
			context.Next()
		}
	}
}

func main() {
	engine := gin.Default()
	//engine.Use(cors.Default()) //  最简单的允许跨域的配置是使用cors.Default(),它默认允许所有跨域请求
	engine.Use(CorsSupport()) //  自定义跨域配置
	engine.GET("/hello", tokenRateLimiter(), HelloHandler)
	engine.Run(":8080")
}

func HelloHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "hello world"})
}

// 自定义跨域配置
func CorsSupport() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"https://foo.com", "https://example.com"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}

	return cors.New(corsConfig)
}
