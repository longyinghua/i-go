package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var mySigningKey = []byte("用于token签名的字符串")

/*自定义token的Claims附加声明*/
type CustomClaims struct {
	jwt.RegisteredClaims        /*jwt包中内嵌的token的声明*/
	Username             string `json:"username"`
}

const TokenExpire = time.Hour * 24         // 定义token的过期时间，常量
var CustomSecret = []byte("用于token签名的字符串") // 定义一个用于签名的字符串

// GenRegisteredClaims
//
//	@Description: 创建token
//	@return string 返回token字符串
//	@return error 错误信息
//
// 创建GenRegisteredClaims函数，使用默认的声明创建jwt,即创建一个token
func GenRegisteredClaims() (string, error) {
	//创建claims，使用jwt的默认claims声明，也就是设置token的额外属性声明
	claims := jwt.RegisteredClaims{
		Issuer:    "longyinghua", /*token的签发人*/
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	//创建token，使用jwt包创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//将token转换成string类型
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		println(err)
		return "", err
	}
	println(tokenString)
	return tokenString, nil
}

// ParseRegisteredClaims
//
//	@Description: 解析token
//	@param tokenString  传入token字符串，用于解析
//	@return bool  token是否有效
//
// 创建ParseRegisteredClaims解析token
func ParseRegisteredClaims(tokenString string) bool {
	//解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		println(err)
		return false
	}

	println(token.Valid)
	return token.Valid
}

// GenToken
//
//	@Description: 创建token，使用自定义的声明Claim，实现创建jwt
//	@return string  返回token字符串
//	@return error  返回错误信息
//
// 创建GenToken函数，使用自定义的声明Claim，实现创建jwt
func GenToken(username string) (string, error) {
	//创建自定义的claims
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Username:         username,
	}

	//创建token，使用jwt包创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//将token转换成string类型
	tokenString, err := token.SignedString(CustomSecret)
	if err != nil {
		println(err)
		return "", err
	}
	println(tokenString)
	return tokenString, nil
}

// ParseToken
//
//	@Description: 解析传入的token字符串
//	@param tokenString   token字符串
//	@return *CustomClaims  返回token的claims
//	@return error         错误信息
//
// 创建ParseToken函数，解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})

	if err != nil {
		println(err)
		return nil, err
	}

	//对token对象中的claim进行类型断言
	claims, ok := token.Claims.(*CustomClaims)
	if ok == true && token.Valid {
		fmt.Println(*claims)
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}

//在gin框架中使用jwt

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthHandle
//
//	@Description: 根据用户提供的用户名和密码，获取token
//	@param context
func AuthHandle(context *gin.Context) {
	//用户发送用户名和密码过来
	var user UserInfo
	err := context.ShouldBind(&user)
	if err != nil {
		println(err)
		context.JSON(
			http.StatusBadRequest,
			gin.H{"msg": "无效参数"},
		)
		return
	}

	//校验用户名和密码是否正确
	if user.Username == "long" && user.Password == "123456" {
		//创建token
		token, err := GenToken(user.Username)
		if err != nil {
			println(err)
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"msg": "服务器错误"},
			)
			return
		}
		//返回token
		context.JSON(
			http.StatusOK,
			gin.H{
				"msg":   "Successful",
				"token": token,
			},
		)
		return
	}
	context.JSON(
		http.StatusOK,
		gin.H{
			"msg":  "用户名或密码错误",
			"user": user.Username,
			"pass": user.Password,
		})
	return
}

// JWTAuthMiddleware 基于jwt的认证中间件
//
//	@Description:  基于jwt的认证中间件，来认证token是否有效
//	@return func(context *gin.Context)
//
// 解析token，用户通过接口获取token之后，后续的请求会携带着token再来请求我们的其他借口
// JWTAuthMiddleware 基于jwt的认证中间件
func JWTAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		//	客户端携带token有三种方式 1.放在请求头 2.放在请求体 3.放在url
		//这里假设token放在了请求头的Authorization中，并使用Bearer 开头
		//	这里的具体实现方式要根据你的实际业务情况决定
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求头中auth为空",
			})
			fmt.Println(authHeader)
			context.Abort()
			return
		}

		//按空格分割请求头中的token，处理错误
		splitN := strings.SplitN(authHeader, " ", 2)
		if !(len(splitN) == 2 && splitN[0] == "Bearer") {
			context.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求头中auth格式有误",
			})
			fmt.Println(authHeader)
			context.Abort()
			return
		}

		//splitN[1]是获取到的请求头中的tokenString，我们使用之前定义好的解析jwt的函数来解析tokenString
		tokenString := splitN[1]
		tokenClaims, err := ParseToken(tokenString)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "无效的token",
			})
			context.Abort()
			return
		}

		//将当前请求的username信息保存到请求的上下文context中
		context.Set("username", tokenClaims.Username)
		context.Next() //后续的处理函数可以通过context.Get("username")来获取当前请求的用户信息
	}
}

// 创建一个/home路由，将通过JWTAuthMiddleware认证的用户来正常响应请求
func homeHandler(context *gin.Context) {
	username := context.MustGet("username")
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"username": username,
			"token":    "lalala",
		},
	})
}

func main() {
	engine := gin.Default()
	//注册一条路由/auth,对外提供获取token的渠道
	engine.POST("/auth", AuthHandle)
	engine.GET("/home", JWTAuthMiddleware(), homeHandler)

	err := engine.Run(":8080")
	if err != nil {
		fmt.Println("项目启动失败，不能监听8080端口")
		return
	}
}
