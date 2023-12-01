package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type SingUpParam struct {
	Age        int    `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Date       string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

var trans ut.Translator

// 在validator验证器中注册翻译器
func InitTrans(locale string) (err error) {
	//	修改gin框架中validator引擎属性，实现自定制
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {

		//注册一个获取json tag的自定义方法
		validator.RegisterTagNameFunc(func(field reflect.StructField) string {
			jsonTagStr := field.Tag.Get("json")

			fmt.Println(jsonTagStr) //  获取json标签
			//typeOf := reflect.TypeOf(jsonTagStr)    //  反射查看变量类型
			//fmt.Println("jsonTagStr type ", typeOf) //  查看变量类型
			//n := strings.SplitN(jsonTagStr, ",", 2)
			//fmt.Println(n)
			//fmt.Println(n[0])

			name := strings.SplitN(jsonTagStr, ",", 2)[0]
			//ss := strings.SplitN(jsonTagStr, ",", 2)

			//fmt.Println("name = ", ss)
			if name == "-" {
				return ""
			}
			return name
		})

		//为SignUpParam注册自定义校验方法
		validator.RegisterStructValidation(SignUpParamStructLevelValidation, SingUpParam{})

		//在验证期中注册自定义校验方法，进行时间字段的自定义校验checkDate

		errs := validator.RegisterValidation("checkDate", customFunc)
		if errs != nil {
			return errs
		}

		zhH := zh.New() //  中文翻译器
		enH := en.New() //  英文翻译器

		uni := ut.New(enH, zhH, enH)

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		//注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(validator, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(validator, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(validator, trans)
		}
		return
	}

	//注意：因为这里会使用到trans实例
	errors := validator.RegisterTranslation(
		"checkDate",
		trans,
		registerTranslator("checkDate", "{0}必须要晚于当前日期"),
		translate,
	)
	if errors != nil {
		return errors
	}

	return
}

// 移除结构体名称前缀
func removeTopStruct(fields map[string]string) map[string]string {
	res := make(map[string]string)
	for field, err := range fields {
		fmt.Println(field[strings.Index(field, ".")+1:]) //  返回field字符串中子字符串。所在的索引位置，加+1表示从.后一位的字符串开始的索引位置，：标识切割到最后
		s := field[strings.Index(field, ".")+1:]
		res[s] = err
		fmt.Printf("res = %v", res)
	}
	return res
}

func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(SingUpParam)
	if su.Password != su.RePassword {
		//	输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}

// 自定义字段级别参数校验
func customFunc(f1 validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", f1.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		err := trans.Add(tag, msg, false)
		if err != nil {
			return err
		}
		return nil
	}
}

// translate自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func main() {
	//init在验证器中注册中文翻译器
	err := InitTrans("zh")
	if err != nil {
		fmt.Printf("init trans failed,err:%v", err)
		return
	}

	engine := gin.Default()
	engine.POST("/singup", signupHandlerFunc)
	err = engine.Run(":8080")
	if err != nil {
		panic(err)
		return
	}
}

func signupHandlerFunc(context *gin.Context) {
	var singupparam SingUpParam
	err := context.ShouldBindJSON(&singupparam)
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}

		//如果是validator.validationErrors类型errors的错误则进行翻译
		//并使用removeTopStruct方法移除结构体名称前缀
		context.JSON(http.StatusOK, gin.H{
			//"msg": errors.Translate(trans),
			"msg": removeTopStruct(errors.Translate(trans)),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg": "singupparm parse success",
	})
	return
}
