package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gorm/dal/model"
	"go-gorm/mysqlcrud"
	"log"
	"net/http"
)

// 删除记录
func BookDelete1(c *gin.Context) {
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

	row, err := mysqlcrud.DeleteBook1(&book)

	c.JSON(http.StatusOK, gin.H{
		"删除记录条数": row,
		"msg":    "success",
		"code":   200,
	})

	return
}
