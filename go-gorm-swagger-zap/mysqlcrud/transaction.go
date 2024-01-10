package mysqlcrud

import (
	"context"
	"errors"
	"go-gorm-swagger-zap/dal"
	"go-gorm-swagger-zap/dal/model"
	"go-gorm-swagger-zap/dal/query"
	"log"
)

// TXCreate1
//
//	@Description:	事务创建,嵌套事务，您可以回滚较大事务内执行的一部分操作
//	@param			book	[]*model.Book
//	@return			error

func TXCreate1(book []*model.Book) error {
	q := query.Use(dal.DB)

	err := q.Transaction(func(tx *query.Query) error {
		//tx.Book.WithContext(context.Background()).Create(&model.Book{Author: "aa", Price: 100})

		tx.Book.WithContext(context.Background()).Create(book...)

		tx.Transaction(func(tx2 *query.Query) error {
			tx2.Book.WithContext(context.Background()).Create(&model.Book{Author: "bb", Price: 200})
			return errors.New("rollback user2") // Rollback 此项
		})

		tx.Transaction(func(tx3 *query.Query) error {
			return tx3.Book.WithContext(context.Background()).Create(&model.Book{Author: "cc", Price: 300})
		})

		return nil
	})

	// Commit user1, user3 // 提交用户 1、用户 3、回滚用户 2

	if err != nil {
		log.Printf("err: %v", err)
		return err
	}

	return nil
}

// TXCreate2
//
//	@Description:	事务创建,手动事务
//	@param			book	*model.Book
//	@return			error error

func TXCreate2(book *model.Book) error {
	u := query.Book
	ctx := context.Background()

	//开始事务
	tx := dal.DB.Begin()

	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	tx.Create(model.Book{
		Author: book.Author,
		Price:  book.Price,
	})

	// 在事务中执行一些 db 操作
	tx.Where(u.WithContext(ctx).Where(u.ID.Eq(book.ID))).Updates(model.Book{Author: book.Author, Price: book.Price})
	// ...

	//遇到错误时回滚事务
	tx.Rollback()

	//否则，提交事务
	tx.Commit()

	return nil
}
