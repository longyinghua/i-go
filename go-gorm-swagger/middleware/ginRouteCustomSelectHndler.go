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

// BookSelect1 查询book书籍信息接口-1
//
//	@Summary		查询book书籍信息
//	@Tags			书籍查询
//	@Description	前端传入book书籍相关信息查询book对象信息的接口，通过author查询
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result{data=[]model.Book}	"查询数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/custom/books1 [post]
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
					"msg":  err.Error(),
					"code": 400,
				})
				return
			}
		}

		books, err := mysqlcrud.CustomSQLSelect1(book.Author)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  err.Error(),
				"code": 400,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": books,
			"msg":  "success",
			"code": 200,
		})
	}
}

// BookSelect2 查询book书籍信息接口-2
//
//	@Summary		查询book书籍信息
//	@Tags			书籍查询
//	@Description	前端传入book书籍相关信息查询book对象信息的接口,通过id查询，返回模型
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result{data=[]model.Book}	"查询数据成功提示"
//	@Failure		400		{object}	ret.Result	"错误提示"
//	@Router			/custom/books2 [post]
func CustomSelect2(ctx *gin.Context) {
	var book model.Book
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"code": 400,
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
				"msg":  err.Error(),
				"code": 400,
			})
			return
		}
	}

	books, err := mysqlcrud.CustomSQLSelect2(book.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
		"msg":  "success",
		"code": 200,
	})

}

// BookSelect3 查询book书籍信息接口-3
//
//	@Summary		查询book书籍信息
//	@Tags			书籍查询
//	@Description	前端传入book书籍相关信息查询book对象信息的接口,通过id查询返回map
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result{data=model.Book}	"查询数据成功提示"
//	@Failure		400		{object}	ret.Result	"错误提示"
//	@Router			/custom/books3 [post]
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
				"msg":  err.Error(),
				"code": 400,
			})
			return
		}
	}

	resultMap, err := mysqlcrud.CustomSQLSelect3(book.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resultMap,
		"msg":  "success",
		"code": 200,
	})
}

// BookSelect4 查询book书籍信息接口-4
//
//	@Summary		查询book书籍信息
//	@Tags			书籍查询
//	@Description	前端传入book书籍相关信息查询book对象信息的接口,通过字段查询返回数据列表集合
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.TFilter	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result{data=[]model.Book}	"查询数据成功提示"
//	@Failure		400		{object}	ret.Result	"错误提示"
//	@Router			/custom/books4 [post]
func CustomSelect4(ctx *gin.Context) {
	var book model.TFilter

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"code": 400,
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
				"msg":  err.Error(),
				"code": 400,
			})
			return
		}
	}

	result, err := query.Book.WithContext(context.Background()).FilterWithColumn(book.Column, book.Value)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"msg":  "success",
		"code": 200,
	})
}

// BookSelect5 查询book书籍信息接口-5
//
//	@Summary		查询book书籍信息
//	@Tags			书籍查询
//	@Description	前端传入book书籍相关信息查询book对象信息的接口,通过书籍字段查询返回数据列表集合
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result{data=[]model.Book}	"查询数据成功提示"
//	@Failure		400		{object}	ret.Result	"错误提示"
//	@Router			/custom/books5 [post]
func CustomSelect5(ctx *gin.Context) {
	var book model.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"code": 400,
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
				"msg":  err.Error(),
				"code": 400,
			})
			return
		}
	}

	result, err := query.Book.WithContext(context.Background()).Search(&book)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"msg":  "success",
		"code": 200,
	})
}

// BookSelect6 查询book书籍信息接口-6
//
//	@Summary		查询book书籍信息
//	@Tags			书籍查询
//	@Description	前端传入book书籍相关信息查询book对象信息的接口,通过字段查询返回数据列表集合
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result{data=[]model.Book}	"查询数据成功提示"
//	@Failure		400		{object}	ret.Result	"错误提示"
//	@Router			/custom/books6 [post]
func CustomSelect6(ctx *gin.Context) {

	var book model.Book

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"code": 400,
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
				"msg":  err.Error(),
				"code": 400,
			})
			return
		}
	}

	result, err := query.Book.WithContext(context.Background()).SearchCustom(&book)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"msg":  "success",
		"code": 200,
	})
}
