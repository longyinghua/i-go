package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strconv"
	_ "test-swagger/docs"
	"test-swagger/ret"
	"time"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

// Hello 测试
//
//	@Summary		测试SayHello
//	@Description	向你说Hello
//	@Tags			测试
//	@Accept			json
//	@Produce		json
//	@Param			who	query		string	true			"人名"
//	@Success		200	{string}	string	"{"msg": "hello	lixd"}"
//	@Failure		400	{string}	string	"{"msg": "who	are	you"}"
//	@Router			/hello [get]
func Hello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}

// Login 登陆
//
//	@Summary		登陆
//	@Tags			登陆注册
//	@Description	登入
//	@Accept			json
//	@Produce		json
//	@Param			user	body		LoginReq					true	"用户名密码"
//	@Success		200		{object}	ret.Result{data=LoginResp}	"token"
//	@Failure		400		{object}	ret.Result					"错误提示"
//	@Router			/login [post]
func Login(c *gin.Context) {
	var m LoginReq
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, ret.Fail("参数错误"))
		return
	}

	if m.Username == "admin" && m.Password == "123456" {
		resp := LoginResp{Token: strconv.Itoa(int(time.Now().Unix()))}
		c.JSON(http.StatusOK, ret.Success(resp))
		return
	}
	c.JSON(http.StatusUnauthorized, ret.Fail("user  or  password  error"))
}

var swagHandler gin.HandlerFunc

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server.
//	@termsOfService	https://lixueduan.com

//	@contact.name	longyinghua
//	@contact.url	https://lixueduan.com
//	@contact.email	542791872@qq.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		10.40.3.9:8080
//	@BasePath	/api/v1

// SwaggerUI: http://localhost:8080/swagger/index.html
func main() {

	e := gin.Default()
	v1 := e.Group("/api/v1", RateLimitMiddleware(1*time.Second, 10))
	{
		v1.GET("/hello", Hello)
		v1.POST("/login", Login)
	}

	if swagHandler != nil {
		e.GET("/swagger/*any", swagHandler)
	}

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerfiles.Handler)
}

// RateLimitMiddleware 限流中间件函数
//
// 参数：
// - fillInterval: 填充间隔时间，即 bucket 的填充时间
// - cap: bucket 的容量
//
// 返回值：
// - gin.HandlerFunc: 限流中间件处理函数
func RateLimitMiddleware(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	//初始化创建新的令牌桶，填充间隔时间为time.Duration秒，容量为cap个数
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(context *gin.Context) {
		// 如果取不到令牌就中断本次请求返回rate limit 。。。
		if bucket.TakeAvailable(1) < 1 {
			context.JSON(http.StatusTooManyRequests, gin.H{"msg": "rate limit ..."})
			context.Abort()
			return
		}
		context.Next()

		var limiter = ratelimit.NewBucketWithQuantum(time.Second, 10, 10)
		log.Printf("limiter.TakeAvailable(1) = %d", limiter.TakeAvailable(1))
	}
}
