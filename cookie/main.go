package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func setCookie(writer http.ResponseWriter, request *http.Request) {
	// 设置cookie
	writer.Header().Set("Set-Cookie", "name=long")
	writer.Header().Set("Content-Type", "application/json")

	// 向客户端发送响应内容
	data := make(map[string]string)
	data["message"] = "Cookie has been set!"
	data["cookie"] = "name=long"
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Write(jsonData)
}

func getCookie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// 从请求头中获取cookie
	cookie, err := request.Cookie("name")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 向客户端发送响应内容
	data := make(map[string]interface{})
	data["message"] = "The value of the cookie is: " + cookie.Value
	data["cookie"] = cookie.String()
	data["status"] = http.StatusOK

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	writer.Write(jsonData)
}

// SetCookie
//
//	@Description: 设置cookie
//	@param writer http.ResponseWriter
//	@param request *http.Request
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	//定义一个cookie切片
	var cookies []*http.Cookie

	c1 := &http.Cookie{Name: "name", Value: "long", Path: "/auth", Expires: time.Now().Add(time.Hour * 1)}
	c2 := &http.Cookie{Name: "pwd", Value: "123456", Path: "/login", Expires: time.Now().Add(time.Hour * 1)}

	cookies = append(cookies, c1, c2)

	//使用http内置函数在响应中写入cookie
	writer.Header().Add("Set-Cookie", cookies[0].String())
	writer.Header().Add("Set-Cookie", cookies[1].String())

	//    json响应输出name:long
	//   设置cookie，cookie为name：long
	//writer.Header().Set("Set-Cookie", "name=long")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("{\"name\":\"long\",\"age\":18}"))
	writer.WriteHeader(200)

}

// GetCookie
//
//	@Description: 从请求头中获取cookie
//	@param writer http.ResponseWriter
//	@param request *http.Request
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	//从请求/auth中获取请求头中的cookie，并在日志中打印出来
	cookie, err := request.Cookie("name")
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("{\"error\":\"cookie not found\"}"))
		return
	}
	// Set the response header content type to JSON
	writer.Header().Set("Content-Type", "application/json")
	// Write the cookie value to the response
	writer.Write([]byte("{\"name\":\"" + cookie.Value + "\"}"))
	// Set the response status code to 200 (OK)
	writer.WriteHeader(200)
	// Log the cookie value
	log.Fatalf("cookie:%v", cookie)
}

func GinCookie(ctx *gin.Context) {
	cookie := &http.Cookie{
		Name:     "name",
		Value:    "long",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
	}
	s, err := ctx.Cookie("name")
	if err != nil {
		fmt.Println(err)

		data := make(map[string]interface{})
		data["message"] = "开始设置cookie"
		data["status"] = http.StatusOK
		data["cookie"] = cookie.String()
		ctx.SetCookie("name", "long", 3600, "/", "localhost", false, true)
		//ctx.Writer.Header().Add("Set-Cookie",cookie.String())
		bytes, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx.Header("Content-Type", "application/json")
		ctx.Writer.Write(bytes)

	}

	fmt.Printf("Cookie value: %s\n", s)

}

func main() {
	//go中使用cookie
	//http.HandleFunc("/set", setCookie)
	//http.HandleFunc("/get", getCookie)
	//http.ListenAndServe(":8080", nil)

	//Gin中使用cookie
	engine := gin.Default()
	engine.GET("/setcookie", GinCookie)
	engine.Run(":8080")
}
