package controller

import (
	"context"
	"gin-gorm-app1/dal/model"
	"gin-gorm-app1/dal/query"
)

// CreateBook1
//
//	@Description:	插入单条数据
//	@param			book	*model.Book	参数为结构体模型，可以根据请求体中的body绑定模型结构体后取具体的字段作为SQL条件进行insert插入
//	@return			error 返回错误信息

func CreateBook1(book *model.Book) error {
	// 创建一个上下文
	ctx := context.Background()
	// 创建一个查询对象
	q := query.Book

	// 插入单条数据
	err := q.WithContext(ctx).Create(book)

	//err := query.Book.WithContext(context.Background()).Create(book) //  插入单条数据

	if err != nil {
		//log.Printf("create book error: %v\n", err)
		return err
	}

	return nil
}
