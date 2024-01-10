package mysqlcrud

import (
	"context"
	"go-gorm-swagger-zap/dal/model"
	"go-gorm-swagger-zap/dal/query"
	"time"
)

// UpdateBook1
//
//	@Description:	更新单列数据
//	@param			book	*model.Book
//	@return			int64 rowsAffected返回更新的行数
//	@return			error error错误信息

func UpdateBook1(book *model.Book) (int64, error) {
	u := query.Book
	ctx := context.Background()
	resultInfo, err := u.WithContext(ctx).Where(u.Author.Eq(book.Author)).Update(u.Price, book.Price)
	if err != nil {
		//log.Printf("update book error: %v", err)
		return 0, err
	}
	rowsAffected := resultInfo.RowsAffected
	return rowsAffected, nil
}

// UpdateBook2
//
//	@Description:	更新多列数据，使用map进行方式
//	@param			book	*model.Book
//	@return			int64 rowsAffected返回更新的行数
//	@return			error error错误信息

func UpdateBook2(book *model.Book) (int64, error) {
	u := query.Book
	ctx := context.Background()
	// 使用 `map` 更新字段
	resultInfo, err := u.WithContext(ctx).Where(u.ID.Lt(10)).Updates(map[string]interface{}{"author": book.Author, "price": book.Price})
	if err != nil {
		//log.Printf("update book error: %v", err)
		return 0, err
	}
	rowsAffected := resultInfo.RowsAffected

	return rowsAffected, nil
}

// UpdateBook3
//
//	@Description:	更新多列数据，使用结构体的方式
//	@param			book	*model.Book
//	@return			int64 rowsAffected返回更新的行数
//	@return			error

func UpdateBook3(book *model.Book) (int64, error) {
	u := query.Book
	ctx := context.Background()
	// 使用 `struct` 更新字段
	resultInfo, err := u.WithContext(ctx).Where(u.ID.Lt(10)).Updates(model.Book{Author: book.Author, Price: book.Price})
	if err != nil {
		//log.Printf("update book error: %v", err)
		return 0, err
	}
	rowsAffected := resultInfo.RowsAffected

	return rowsAffected, nil
}

// UpdateBook4
//
//	@Description:	更新多列数据，更新选定字段,使用结构体的方式
//	@param			book	*model.Book
//	@return			int64 rowsAffected返回更新的行数
//	@return			error error错误信息

func UpdateBook4(book *model.Book) (int64, error) {
	u := query.Book
	ctx := context.Background()
	//更新选定字段，使用Select
	resultInfo, err := u.WithContext(ctx).Select(u.Author, u.Price).Where(u.ID.Eq(book.ID)).Updates(model.Book{
		ID:          0,
		Title:       "",
		Author:      book.Author,
		Price:       book.Price,
		PublishDate: time.Time{},
		CreateAt:    time.Time{},
		UpdateAt:    time.Time{},
	})
	if err != nil {
		//log.Printf("update book error: %v", err)
		return 0, err
	}
	rowsAffected := resultInfo.RowsAffected
	return rowsAffected, nil
}
