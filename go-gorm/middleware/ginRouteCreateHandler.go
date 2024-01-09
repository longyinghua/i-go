package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gorm/dal/model"
	"go-gorm/mysqlcrud"
	"log"
	"net/http"
)

func BookCreate1() gin.HandlerFunc {
	return func(c *gin.Context) {

		var book model.Book

		err := c.ShouldBindJSON(&book)
		if err != nil {
			log.Printf("book param bind json err:%v", err)
		}

		//  参数校验
		if err != nil {
			//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
			_, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					//不是validator.validationErrors类型errors的错误直接返回错误信息
					"msg":  err.Error(),
					"code": 500,
				})
				return
			}
		}

		err = mysqlcrud.CreateBook1(&book) //  插入数据成功返回true，失败返回false

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("保存数据失败:%v", err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  fmt.Sprintf("保存数据成功"),
			})
		}

	}
}

func BookCreate2(c *gin.Context) {

	var book model.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Printf("book param bind json err:%v", err)
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg":  err.Error(),
				"code": 500,
			})
			return
		}
	}

	err = mysqlcrud.CreateBook1(&book) //  插入数据成功返回true，失败返回false

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  fmt.Sprintf("保存数据失败:%v", err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  fmt.Sprintf("保存数据成功"),
		})
	}

}

func BookCreate3(c *gin.Context) {

	// 解析JSON请求体
	var requestPayload struct {
		Books []*model.Book `json:"books"`
	}

	err := c.ShouldBindJSON(&requestPayload)

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"error": err.Error(),
				"code":  500,
			})
			return
		}
	}

	err = mysqlcrud.CreateBook2(requestPayload.Books) //  插入数据成功返回true，失败返回false

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  fmt.Sprintf("保存数据失败:%v", err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  fmt.Sprintf("保存数据成功"),
		})
	}

}

//apifox接口传参示例
//{
//"books": [
//{
//"title": "<<C语言之路>>",
//"author": "wawa",
//"price": 100
//},
//{
//"title": "<<C++语言之路>>",
//"author": "wlalaawa",
//"price": 200
//},
//{
//"title": "<<Kubernetes进阶指南>>",
//"author": "kakak",
//"price": 300
//}
//],
//"hello": "world"
//}

func BookCreate4(c *gin.Context) {

	//引用全局自定义的模型
	var requestPlayload model.RequestPayloadBook

	//解析JSON请求体，将json请求体绑定到requestPlayload中，api传入的books的key对应requestPlayload.Books的key名，books的value为列表对应requestPlayload.Books的value，【】*Book
	err := c.ShouldBindJSON(&requestPlayload)
	//  注意：参数绑定时，一定要包含对应的key，key对应模型中对应的字段名，如果前端传入的key名和模型中字段名不一致，则不会绑定到对应模型的字段中

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"error": err.Error(),
				"code":  500,
			})
			return
		}
	}

	err = mysqlcrud.CreateBook2(requestPlayload.Books) //  插入数据成功返回true，失败返回false

	//BookBody := requestPlayload.Books
	//err = mysqlcrud.CreateBook2(BookBody) //  插入数据成功返回true，失败返回false

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  fmt.Sprintf("保存数据失败:%v", err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  fmt.Sprintf("保存数据成功"),
		})
	}

}
