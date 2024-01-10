package mysqlcrud

import (
	"context"
	"fmt"
	"go-gorm-swagger-zap/dal/model"
	"go-gorm-swagger-zap/dal/query"
	"log"
)

// SelectBook1
//
//	@Description:	查询表中的单条数据，如果没有找到记录，则返回 ErrRecordNotFound 错误，只有正向或逆向排序的一条数据
//	@return			*model.Book 返回查询到的单条数据

func SelectBook1() *model.Book {
	//	查询表中数据
	// 当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误
	book, err := query.Book.WithContext(context.Background()).First() //  单条查询通过逐渐id正向排序取到第一条数据 Get the first record ordered by primary key
	//book, err := query.Book.WithContext(context.Background()).Last() //  单条查询通过逐渐id逆向排序取到第一条数据 Get the first record ordered by primary key
	//book, err := query.Book.WithContext(context.Background()).Take() //  单条查询，只取一条数据，如果没有找到记录，则返回 ErrRecordNotFound 错误
	if err != nil {
		fmt.Printf("query book error: %v\n", err)
		return nil
	}
	fmt.Printf("query book result: %v\n", book)

	return book
}

// SelectBook2
//
//	@Description:	查询表中的所有数据
//	@return			[]*model.Book 返回查询到的所有数据

func SelectBook2() []*model.Book {
	books, err := query.Book.WithContext(context.Background()).Find() //   SELECT * FROM `book`
	if err != nil {
		fmt.Printf("query book error: %v\n", err)
		return nil
	}

	return books
}

// SelectBook3
//
//	@Description:	查询表中的所有数据
//	@param			id	...int64	根据传入的id查询数据
//	@return			[]*model.Book

func SelectBook3(id ...int64) []*model.Book {
	books, err := query.Book.WithContext(context.Background()).Where(query.Book.ID.In(id...)).Find()
	if err != nil {
		log.Printf("query book error: %v\n", err)
		return nil
	}

	return books
}

// SelectBook4
//
//	@Description:	查询表中的符合条件的所有数据
//	@param			book	*model.Book	参数为结构体模型，可以根据请求体中的body绑定模型结构体后取具体的字段作为SQL条件进行查询
//	@return			[]*model.Book 返回查询到的所有数据，结果集

func SelectBook4(book *model.Book) []*model.Book {
	u := query.Book
	ctx := context.Background()
	books, err := u.WithContext(ctx).Where(u.ID.Gt(book.ID), u.Author.Neq(book.Author)).Find() //  book.ID,book.Author为结构体中的字段，可以根据请求体中的body绑定模型结构体后取具体的字段作为SQL条件进行查询
	if err != nil {
		fmt.Printf("query book error: %v\n", err)
		return nil
	}

	return books
}

// SelectBook5
//
//	@Description:	查询表中的符合条件的所有数据
//	@param			book	*model.Book	参数为结构体模型，可以根据请求体中的body绑定模型结构体后取具体的字段作为SQL条件进行查询
//	@return			[]*model.Book 返回查询到的所有数据，结果集

func SelectBook5(book *model.Book) []*model.Book {
	u := query.Book
	ctx := context.Background()
	books, err := u.WithContext(ctx).Where(u.ID.Gt(book.ID), u.Author.Neq(book.Author)).Find()
	if err != nil {
		fmt.Printf("query book error: %v\n", err)
		return nil
	}

	return books
}
