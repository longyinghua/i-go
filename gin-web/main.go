package main

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// 获取url当中的路径参数，即/user/jack/get
func Getapi() {
	//使用gin创建路由
	r := gin.Default()
	//绑定路由规则，执行的函数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截图
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//监听8888端口
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("服务器启动成功")
	fmt.Printf("服务器启动成功")

}

// 获取url当中的请求参数，即?name=jack&age=18&key=value
func Geturl() {
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "jack") //  url请求参数当中指定的key，这个key名字为name时，输出其参数对应的值，如果key不存在，返回空值
		ctx.String(http.StatusOK, name)
	})
	//监听8888端口
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func postform() {
	//获取表单参数
	r := gin.Default()
	r.POST("/form", func(ctx *gin.Context) {
		//获取表单参数
		types := ctx.DefaultPostForm("type", "post")
		username := ctx.PostForm("username")
		password := ctx.PostForm("userpassword")
		ctx.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run("0.0.0.0:8888")
}

func upfile() {
	//上传文件
	r := gin.Default()
	//设置文件最大上传尺寸为8M
	r.MaxMultipartMemory = 8 << 20
	//限制只能上传jpg,png类型的文件

	r.POST("/upload", func(ctx *gin.Context) {
		//获取文件
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
		}
		ctx.SaveUploadedFile(file, file.Filename)
		ctx.String(http.StatusOK, fmt.Sprintf("file:%s", file.Filename))
	})
	r.Run("0.0.0.0:8888")
}

func upfilem() {
	//gin创建路由
	r := gin.Default()
	//设置文件最大上付尽数为8M
	r.MaxMultipartMemory = 8 << 20
	//upload接口上传多个文件
	r.POST("/upload", func(ctx *gin.Context) {
		//获取所有上传文件
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
			return
		}
		files := form.File["files"]
		for _, file := range files {
			//获取文件名，上传
			err := ctx.SaveUploadedFile(file, file.Filename)
			if err != nil {
				ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err:%s", err.Error()))
				return
			}
			ctx.String(http.StatusOK, fmt.Sprintf("file:%s\n", file.Filename))
		}
		//输出总计上传了多少个文件
		ctx.String(http.StatusOK, fmt.Sprintf("total file num:%d\n", len(files)))
	})
	r.Run("0.0.0.0:8888")
}

func routeGroup() {
	//1.创建路由
	r := gin.Default()
	//路由组1,处理Get请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	//创建组2，处理Post请求
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	//运行在本机的8888端口
	r.Run("0.0.0.0:8888")
}

func submit(context *gin.Context) {
	context.String(http.StatusOK, "submit")
	return
}

func login(context *gin.Context) {
	context.String(http.StatusOK, "login")
	return
}

// 定义接收数据的结构体
type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func jsonPar() {
	//	1.创建路由
	r := gin.Default()
	//Json绑定
	r.POST("/loginJSON", func(context *gin.Context) {
		//声明接收的变量
		var json Login
		//将request的body中的数据，自动按照json格式解析到结构体中
		if err := context.ShouldBindJSON(&json); err != nil {
			//返回错误信息
			//gin.H封装了生成json数据的工具
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		//返回成功信息
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	//进行監听
	r.Run("0.0.0.0:8888")
}

// 表单数据解析和绑定
func formPar() {
	type Login struct {
		//binding：“required”修饰的字段，若接收为空值，则报错，是必须字段
		User     string `form:"username1" json:"user" uri:"user" xml:"user" binding:"required"`
		Password string `form:"password1" json:"password" uri:"password" xml:"password" binding:"required"`
	}

	r := gin.Default()
	r.POST("/loginForm", func(context *gin.Context) {
		var form Login
		//Parse the binding form format
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if form.User != "root" || form.Password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		//返回成功信息
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run("0.0.0.0:8888")
}

func urlPar() {
	//url数据解析和绑定
	//定义一个名为Login的结构体，字段有User，Password
	type Login struct {
		User string `json:"user" xml:"user" form:"user" uri:"user" binding:"required"`
		Pass string `json:"pass" xml:"pass" form:"pass" uri:"pass" binding:"required"`
	}
	r := gin.Default()
	r.GET("/login/:user/:pass", func(context *gin.Context) {
		var uri Login
		//Parse the binding uri format
		if err := context.ShouldBindUri(&uri); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if uri.User != "root" || uri.Pass != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		//返回成功信息
		context.JSON(http.StatusOK, gin.H{"status": "200"})
		return
	})
	r.Run("0.0.0.0:8888")
}

func responseDate() {
	//多重响应方式
	r := gin.Default()
	r.GET("/responsejson", func(context *gin.Context) {
		//返回json样式
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	//返回xml样式
	r.GET("/responsexml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"status": "200"})
	})
	//返回text样式
	r.GET("/responsetext", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	//返回html样式
	r.GET("/responsehtml", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "hello world"})
	})
	//返回struct样式
	r.GET("/responsestruct", func(context *gin.Context) {
		var msg struct {
			Name    string
			Age     int
			Mumber  int
			Marryed bool
		}
		msg.Name = "小明"
		msg.Age = 18
		msg.Mumber = 123456
		msg.Marryed = false
		context.JSON(http.StatusOK, msg)
	})
	//返回map样式
	r.GET("/responsemap", func(context *gin.Context) {
		msg := gin.H{
			"name":    "小明",
			"age":     18,
			"mumber":  123456,
			"marryed": false,
		}
		context.JSON(http.StatusOK, msg)
	})
	//返回slice样式
	r.GET("/responseslice", func(context *gin.Context) {
		msg := []gin.H{
			{
				"name":    "小明",
				"age":     18,
				"mumber":  123456,
				"marryed": false,
			},
		}
		context.JSON(http.StatusOK, msg)
	})

	//返回array样式
	r.GET("/responsearray", func(context *gin.Context) {
		msg := gin.H{
			"name":    "小明",
			"age":     18,
			"mumber":  123456,
			"marryed": false,
		}
		context.JSON(http.StatusOK, msg)

	})
	//返回yaml格式
	r.GET("/responseyaml", func(context *gin.Context) {
		msg := gin.H{
			"name":    "小明",
			"age":     18,
			"mumber":  123456,
			"marryed": false,
		}
		context.YAML(http.StatusOK, msg)

	})
	//构建一个传输格式
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	//监听在0.0.0.0:8888
	r.Run("0.0.0.0:8888")
}

func redirect() {
	engine := gin.Default()
	engine.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	engine.Run(":8888")
}

func syncbu() {
	//	创建路由
	r := gin.Default()
	//	1.异步
	//	get 请求路径/long_async 路由处理函数
	r.GET("/long_async", func(context *gin.Context) {
		//创建一个副本
		c := context.Copy()
		go func() {
			//延时2秒
			time.Sleep(time.Second * 2)
			//返回结果
			log.Println("Done! in path " + c.Request.URL.Path)
		}()
	})
	//2.同步
	//	get请求/async_async 路由处理
	r.GET("/sync_async", func(context *gin.Context) {
		//延时2秒
		time.Sleep(time.Second * 2)
		//返回结果
		log.Println("Done! in path " + context.Request.URL.Path)
	})

	//监听8888端口
	r.Run(":8888")
}

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		//设置变量到Context的key中，可以通过Get获取
		context.Set("request", "中间件")
		status := context.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func getRoute() {
	//1.创建路由
	engine := gin.Default()
	engine.Use(MiddleWare())
	//规范代码
	{
		engine.GET("/middleware", func(context *gin.Context) {
			//取设置的变量
			value, exists := context.Get("request")
			if exists {
				fmt.Println("request:", value)
				//返回JSON
				context.JSON(http.StatusOK, gin.H{"request": value})
			}
		})
	}

	engine.Run(":8888")
}

func getRoute1() {
	engine := gin.Default()
	//跟路名
	engine.GET("/group", MiddleWare(), func(context *gin.Context) {
		//从局部中间件middleware中取出request的值
		value, exists := context.Get("request")
		if exists {
			fmt.Println("request:", value)
		}
		context.JSON(http.StatusOK, gin.H{"request": value})
	})

	engine.Run(":8888")
}

func myTime(context *gin.Context) {
	now := time.Now()
	context.Next()
	spend := time.Since(now)
	fmt.Println("程序用时:", spend)
}

func api1(ctx *gin.Context) {
	time.Sleep(time.Second * 2)
}

func api2(ctx *gin.Context) {
	time.Sleep(time.Second * 2)
}

func getRoute2() {
	engine := gin.Default()
	engine.Use(myTime)
	//创建api分组
	api := engine.Group("/api")
	{
		//api.GET("/article", api1)
		//api.GET("/user", api2)

		api.GET("/user", func(context *gin.Context) {
			//返回JSON
			context.JSON(http.StatusOK, gin.H{"request": "/user"})
		})

		api.GET("/article", func(context *gin.Context) {
			//返回JSON
			context.JSON(http.StatusOK, gin.H{"request": "/article"})
		})

	}
	engine.Run(":8888")
}

func setcookie() {
	engine := gin.Default()
	engine.GET("/cookie", func(context *gin.Context) {
		//获取客户端是否携带cookie
		cookie, err := context.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"

		}
		//设置cookie
		context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		fmt.Printf("cookie的值是：%s", cookie)
	})
	engine.Run(":8888")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		cookie, err := c.Cookie("abc")
		if err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": err, "abc": cookie})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort() //  终止上下文后，后续的处理函数将不会被执行
		return
	}
}

func loginCookie() {
	engine := gin.Default()
	engine.GET("/login", func(context *gin.Context) {
		//设置cookie
		context.SetCookie("abc", "123", 3600, "/", "localhost", false, true)
		//返回信息
		context.String(http.StatusOK, "login Success!")
	})
	engine.GET("/home", AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	engine.Run(":8888")
}

var store *sessions.CookieStore

func init() {
	//初始化会话存储
	store = sessions.NewCookieStore([]byte("secret"))

}

func sessionget() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/logout", logoutHandler)

	http.Handle("/", router)
	http.ListenAndServe(":8888", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//获取会话
	session, _ := store.Get(r, "my-session")

	//检查用户是否已登录
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//显示欢迎信息
	fmt.Fprintln(w, "Hello, world!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//获取会话
	session, _ := store.Get(r, "my-session")

	//设置会话值，表示用户已登录
	session.Values["authenticated"] = true
	session.Save(r, w)
	//重定向到主页
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	//获取会话
	session, _ := store.Get(r, "my-session")

	//删除会话值，表示用户已注销
	delete(session.Values, "authenticated")
	session.Save(r, w)

	//重定向到主页
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logoutput() {
	gin.DisableConsoleColor() //  禁用控制台颜色输出

	//输出到日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	engine.Run(":8888")
}

// 中间件，处理session
func Session(KeyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(KeyPairs, store)
}

func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "sexxion"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge:   sessionMaxAge,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return store
}

func Captcha(ctx *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[0]
	}
	if len(length) == 3 {
		h = length[2]
	}

	captchaId := captcha.NewLen(1)
	session := sessions.Default(ctx)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Server(ctx.Writer, ctx.Request, captchaId, ".png", "zh", false, w, h)
}

func CaptchaVerify(ctx *gin.Context, code string) bool {
	session := sessions.Default(ctx)
	if captchaId := session.get("captcha"); captchaId != nil {
		sessions.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Server(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch expr {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return fmt.Errorf("wrong captcha type")
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func dosms() {
	engine := gin.Default()
	engine.LoadHTMLFiles("*.html")
	engine.Use(Session("topgoer"))
	engine.GET("/captcha", func(context *gin.Context) {
		Captcha(context, 4, 100, 40)
	})

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	engine.GET("/captcha/verify/:value", func(context *gin.Context) {
		value := context.Param("value")
		if CaptchaVerify(context, value) {
			context.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			context.JSON(http.StatusOK, gin.H{"status": "fail"})
		}
	})
	engine.Run(":8888")
}

func main() {
	//Getapi()
	//Geturl()
	//postform()
	//upfile()
	//upfilem()
	//routeGroup()
	//jsonPar()
	//formPar()
	//urlPar()
	//responseDate()
	//redirect()
	//syncbu()
	//getRoute()
	//getRoute1()
	//getRoute2()
	//setcookie()
	//loginCookie()
	//sessionget()
	//logoutput()
	dosms()
}
