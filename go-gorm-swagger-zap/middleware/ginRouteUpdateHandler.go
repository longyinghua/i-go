package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gorm-swagger-zap/dal/model"
	"go-gorm-swagger-zap/logger"
	"go-gorm-swagger-zap/mysqlcrud"
	"go.uber.org/zap"
	"net/http"
)

// 更新单列
// BookUpdate1 新增更新book书籍信息接口-1
//
//	@Summary		更新book书籍信息
//	@Tags			书籍更新
//	@Description	前端传入book书籍相关信息直接更新book对象信息的接口，更新单列字段
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1	"更新数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/update/books1 [post]
func BookUpdate1(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		//log.Printf("bind json error: %v\n", err)
		logger.Logger.Error("book param bind json error", zap.String("error", err.Error()))
		//zap.L().Error("book param bind json error",zap.String("error",err.Error()))
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg":  err.Error(),
				"code": 400,
			})
			logger.Logger.Error("参数校验失败", zap.String("error", err.Error()))
			//zap.L().Error("参数校验失败",zap.String("error",err.Error()))

			return
		}
	}

	row, err := mysqlcrud.UpdateBook1(&book)

	c.JSON(http.StatusOK, gin.H{
		"更新记录条数": row,
		"msg":    "success",
		"code":   200,
	})
	logger.SugaredLogger.Infof("更新数据成功,更新记录条数：%d", row)
	//zap.L().Sugar().Infof("更新数据成功,更新记录条数：%d",row)
	return
}

// 更新多列
// BookUpdate2 新增更新book书籍信息接口-2
//
//	@Summary		更新book书籍信息
//	@Tags			书籍更新
//	@Description	前端传入book书籍相关信息直接更新book对象信息的接口，更新多列字段，map写法
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1	"更新数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/update/books2 [post]
func BookUpdate2(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		//log.Printf("bind json error: %v\n", err)
		logger.Logger.Error("book param bind json error", zap.String("error", err.Error()))
		//zap.L().Error("book param bind json error",zap.String("error",err.Error()))
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg":  err.Error(),
				"code": 400,
			})
			logger.Logger.Error("参数校验失败", zap.String("error", err.Error()))
			//zap.L().Error("参数校验失败",zap.String("error",err.Error()))
			return
		}
	}

	row, err := mysqlcrud.UpdateBook2(&book)

	c.JSON(http.StatusOK, gin.H{
		"更新记录数": row,
		"msg":   "success",
		"code":  200,
	})
	logger.SugaredLogger.Infof("更新数据成功,更新记录条数：%d", row)
	//zap.L().Sugar().Infof("更新数据成功,更新记录条数：%d",row)

	return
}

// 更新多列
// BookUpdate3 新增更新book书籍信息接口-3
//
//	@Summary		更新book书籍信息
//	@Tags			书籍更新
//	@Description	前端传入book书籍相关信息直接更新book对象信息的接口，更新多列字段，模型写法
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1	"更新数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/update/books3 [post]
func BookUpdate3(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		//log.Printf("bind json error: %v\n", err)
		logger.Logger.Error("book param bind json error", zap.String("error", err.Error()))
		//zap.L().Error("book param bind json error",zap.String("error",err.Error()))
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			logger.Logger.Error("参数校验失败", zap.String("error", err.Error()))
			//zap.L().Error("参数校验失败",zap.String("error",err.Error()))
			return
		}
	}

	row, err := mysqlcrud.UpdateBook3(&book)

	c.JSON(http.StatusOK, gin.H{
		"更新记录数": row,
		"msg":   "success",
		"code":  200,
	})
	logger.SugaredLogger.Infof("更新数据成功,更新记录条数：%d", row)
	//zap.L().Sugar().Infof("更新数据成功,更新记录条数：%d",row)

	return
}

// 更新多列，更新选定字段
// BookUpdate4 新增更新book书籍信息接口-4
//
//	@Summary		更新book书籍信息
//	@Tags			书籍更新
//	@Description	前端传入book书籍相关信息直接更新book对象信息的接口，更新多列，更新选定字段
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result1	"更新数据成功提示"
//	@Failure		400		{object}	ret.Result1	"错误提示"
//	@Router			/update/books4 [post]
func BookUpdate4(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		//log.Printf("bind json error: %v\n", err)
		logger.Logger.Error("book param bind json error", zap.String("error", err.Error()))
		//zap.L().Error("book param bind json error",zap.String("error",err.Error()))
		return
	}

	//  参数校验
	if err != nil {
		//获取validator.validationErrors类型errors，也就是参数不符合tag标签的校验类型错误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				//不是validator.validationErrors类型errors的错误直接返回错误信息
				"msg": err.Error(),
			})
			logger.Logger.Error("参数校验失败", zap.String("error", err.Error()))
			//zap.L().Error("参数校验失败",zap.String("error",err.Error()))
			return
		}
	}

	row, err := mysqlcrud.UpdateBook4(&book)

	c.JSON(http.StatusOK, gin.H{
		"更新记录数": row,
		"msg":   "success",
		"code":  200,
	})
	logger.SugaredLogger.Infof("更新数据成功,更新记录条数：%d", row)
	//zap.L().Sugar().Infof("更新数据成功,更新记录条数：%d",row)

	return
}
