package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gorm/dal/model"
	"go-gorm/dal/query"
	"go-gorm/mysqlcrud"
	"log"
	"net/http"
)

func CustomSelect1() gin.HandlerFunc {
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

		books, err := mysqlcrud.CustomSQLSelect1(book.Author)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"book": books,
			"msg":  "success",
			"code": 200,
		})
	}
}

func CustomSelect2(ctx *gin.Context) {
	var book model.Book
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	books, err := mysqlcrud.CustomSQLSelect2(book.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": books,
		"msg":  "success",
		"code": 200,
	})

}

func CustomSelect3(ctx *gin.Context) {
	var book model.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	resultMap, err := mysqlcrud.CustomSQLSelect3(book.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": resultMap,
		"msg":  "success",
		"code": 200,
	})
}

func CustomSelect4(ctx *gin.Context) {
	var book model.TFilter

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	result, err := query.Book.WithContext(context.Background()).FilterWithColumn(book.Column, book.Value)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": result,
		"msg":  "success",
		"code": 200,
	})
}

func CustomSelect5(ctx *gin.Context) {
	var book model.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	result, err := query.Book.WithContext(context.Background()).Search(&book)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": result,
		"msg":  "success",
		"code": 200,
	})
}

func CustomSelect6(ctx *gin.Context) {

	var book model.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			return
		}
	}

	result, err := query.Book.WithContext(context.Background()).SearchCustom(&book)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": result,
		"msg":  "success",
		"code": 200,
	})
}
