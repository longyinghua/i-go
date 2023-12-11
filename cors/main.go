package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Student struct {
	ID   int64
	Name string
	Sex  string
	Age  int64
}

// 连接数据库
func init() {
	dns := "root:dfzy_12345@tcp(192.168.2.250:3306)/test?charset=utf8&parseTime=True&loc=Local"
	dbcon, err := sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
		return
	}
	db = dbcon
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
}

// 获取所有学生信息(数据自己事先插入)
func GetAllStudent() []Student {
	sqlStr := "SELECT * FROM student"
	students := make([]Student, 0)
	rows, _ := db.Query(sqlStr) //  执行sql查询语句将结果存储到rows变量中
	student := Student{}
	for rows.Next() { //  rows.next判断是否有下一行数据，如果有返回true则执行rows.scan将数据存储到student变量中，否则跳出循环
		rows.Scan(&student.ID, &student.Name, &student.Sex, &student.Age) //  调用rows.scan将数据存储到student变量中
		students = append(students, student)
	}
	defer rows.Close()
	return students
}

func GetAllStudentInfo(w http.ResponseWriter, r *http.Request) {
	students := GetAllStudent()
	for _, student := range students {
		fmt.Println(student)
	}
	resp := make(map[string]interface{})
	resp["msg"] = "成功"
	resp["code"] = "200"
	resp["data"] = students
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(jsonResp)
}

func GetAllStudentInfo2() gin.HandlerFunc {
	return func(context *gin.Context) {
		//context.JSON(http.StatusOK, student)
		student := GetAllStudent()
		//for _, students := range student {
		//	fmt.Println(students)
		//}
		context.JSON(
			http.StatusOK,
			gin.H{
				"code": http.StatusOK,
				"msg":  "Success",
				"data": student,
			},
		)
	}
}

func CossSupport() gin.HandlerFunc {
	config := cors.Config{
		AllowAllOrigins:  false,                                                       //  是否允许所有域名跨域请求,ture则允许所有域名跨域请求
		AllowOrigins:     []string{"https://foo.com", "https://example.com"},          //  是否允许所有域名跨域请求,*表示允许所有域名跨域请求
		AllowCredentials: true,                                                        //  是否允许请求带有验证信息
		AllowHeaders:     []string{"Orign"},                                           //  允许跨域请求的请求头
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "OPTIONS", "DELETE"}, //  允许跨域请求的方法
		ExposeHeaders:    []string{"Content-Length"},                                  //  允许跨域请求的响应头
		AllowOriginFunc: func(origin string) bool { //  是否允许跨域请求的来源，如果设置了AllowAllOrigins为false，则此函数才会生效,如果同时存在AllowOrigins，则AllowOriginFunc优先级高，AllowOrigns不生效
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour, //    跨域请求的有效时间
	}
	return cors.New(config)
}

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method               //  获取请求方法
		origin := context.Request.Header.Get("Origin") // 获取客户端的origin
		if origin != "" {
			//接收客户端发送的origin （重要！）
			context.Writer.Header().Set("Access-Control-Allow-Origin", origin) //  根据客户端实际情况设置，从客户端请求的origin中获取，如果没有则设置为*
			//context.Writer.Header().Set("Access-Control-Allow-Origin", "*") //  设置为*表示允许所有的域名跨域请求，可将*改为指定的域名
			//服务器支持的所有跨域请求的方法
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			context.Header("Access-Control-Allow-Headers", "Authorization, content-type, Content-Length, X-CSRF-Token, Token,session,Access-Control-Allow-Headers,account")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			context.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			context.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			//context.JSON(http.StatusOK, "ok!")
			context.AbortWithStatus(http.StatusNoContent) // 如果客户端发送的是 OPTIONS 请求，那么中间件会中止响应并返回 HTTP 状态码 204
		}

		context.Next()
	}
}

func main() {
	//普通http请求响应
	//http.HandleFunc("/api/students", GetAllStudentInfo)
	//http.ListenAndSer}ve(":8080", nil)

	//gin框架http请求响应
	engine := gin.Default()
	//跨域中间件，必须在路由之前
	//engine.Use(Cors()) //  自定义跨域中间件
	engine.Use(CossSupport()) //  cors框架跨域中间件

	//engine.GET("/api/students", GetAllStudentInfo2())
	engine.GET("/api/students", GetAllStudentInfo2())
	engine.Run(":8080")
}
