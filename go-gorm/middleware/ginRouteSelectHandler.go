package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gorm/dal/model"
	"go-gorm/mysqlcrud"
	"log"
	"net/http"
)

func BookInfo1() gin.HandlerFunc {
	return func(c *gin.Context) {
		book1 := mysqlcrud.SelectBook1()
		c.JSON(http.StatusOK, gin.H{
			"book": book1,
			"msg":  "success",
			"code": 200,
		})
	}
}

func BookInfo2() gin.HandlerFunc {
	return func(c *gin.Context) {
		book2 := mysqlcrud.SelectBook2()
		c.JSON(http.StatusOK, gin.H{
			"book": book2,
			"msg":  "success",
			"code": 200,
		})
	}
}

func BookInfo3() gin.HandlerFunc {
	return func(c *gin.Context) {

		var book model.Book

		err := c.ShouldBindJSON(&book)
		if err != nil {
			log.Printf("book param bind json err:%v", err)
			return
		}

		//  参数校验
		if err != nil {
			//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
			_, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					//不是validator.validationErrors类型errors的错误直接返回错误信息
					"msg": err.Error(),
				})
				return
			}
		}

		book3 := mysqlcrud.SelectBook3(1, 2, 3, 4, 5, 6) // SQL中的 id in 查询
		c.JSON(http.StatusOK, gin.H{
			"book": book3,
			"msg":  "success",
			"code": 200,
		})
	}
}

func BookInfo4() gin.HandlerFunc {
	return func(c *gin.Context) {

		var book model.Book

		err := c.ShouldBindJSON(&book)
		if err != nil {
			log.Printf("book param bind json err:%v", err)
			return
		}

		//  参数校验
		if err != nil {
			//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
			_, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					//不是validator.validationErrors类型errors的错误直接返回错误信息
					"msg": err.Error(),
				})
				return
			}
		}

		book4 := mysqlcrud.SelectBook4(&book)
		c.JSON(http.StatusOK, gin.H{
			"book": book4,
			"msg":  "success",
			"code": 200,
		})
	}
}

func BookInfo5(c *gin.Context) {
	var book model.Book

	//  根据json格式的数据绑定到book对象中
	err := c.ShouldBindJSON(&book)

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	//接口post请求传递json参数绑定到book对象中后，执行SQL查询
	book5 := mysqlcrud.SelectBook5(&book) //  执行SQL查询后返回结果集

	//  将查询结果返回给客户端
	c.JSON(http.StatusOK, gin.H{
		"book": book5,
		"msg":  "success",
		"code": 200,
	})

	return
}

func BookInfo6(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Printf("bind json error: %v\n", err)
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	book4 := mysqlcrud.SelectBook4(&book)

	c.JSON(http.StatusOK, gin.H{
		"book": book4,
		"msg":  "success",
		"code": 200,
	})

	return
}

// curl --location --request POST 'http://10.40.3.9:8080/api/v1/books5' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 10.40.3.9:8080' \
// --header 'Connection: keep-alive' \
// --data-raw '{
// "id": 3,
// "author": "qimi"
// }'
