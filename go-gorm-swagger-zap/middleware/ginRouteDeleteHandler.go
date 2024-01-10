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

// 删除记录
// BookDelete1 删除book书籍信息接口-1
//
//	@Summary		删除book书籍信息
//	@Tags			书籍删除
//	@Description	前端传入book书籍相关信息删除book对象信息的接口
//	@Accept			application/json
//	@Produce		application/json
//	@Param			book	body		model.Book	true	"book书籍的title，author，price相关信息"
//	@Success		200		{object}	ret.Result	"删除数据成功提示"
//	@Failure		400		{object}	ret.Result	"错误提示"
//	@Router			/delete/books1 [post]
func BookDelete1(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		//log.Printf("bind json error: %v\n", err)
		logger.Logger.Error("book param bind json error", zap.String("eroor", err.Error()))
		//zap.L().Error("book param bind json error", zap.String("eroor", err.Error()))
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
			logger.Logger.Warn("book param validate error", zap.String("eroor", err.Error()))
			//zap.L().Warn("book param validate error", zap.String("eroor", err.Error()))
			return
		}
	}

	row, err := mysqlcrud.DeleteBook1(&book)

	c.JSON(http.StatusOK, gin.H{
		"删除记录条数": row,
		"msg":    "success",
		"code":   200,
	})
	logger.SugaredLogger.Infof("delete success num %d", row)

	return
}
