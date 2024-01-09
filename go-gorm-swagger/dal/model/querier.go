package model

import "gorm.io/gen"

//通过添加注释生成自定义方法

type Querier interface {
	//切记注意：sql语句的注释上方一定不要写其他相关的注释，否则会被作为sql语句解析，导致生成的代码错误，注释可以写在上方

	//	Select * from @@table where id=@id
	GetByID(id int64) (gen.T, error) //返回结构体和error，
	// 示例结果：{3 <<Go语言之路>> 七米 100 2024-01-02 09:47:23 +0000 UTC 2024-01-02 09:47:15 +0000 UTC 2024-01-02 09:47:20 +0000 UTC}

	//Select * from @@table where id=@id
	GetByIDReturnMap(id int64) (gen.M, error) //GetByIDReturnMap 根据id查询返回map
	// 示例结果： map[author:七米 create_at:2024-01-02 09:47:15 +0000 UTC id:3 price:100 publish_date:2024-01-02 09:47:23 +0000 UTC title:<<Go语言之路>> update_at:2024-01-02 09:47:20 +0000 UTC]

	//select * from @@table where author=@author
	GetBooksByAuthor(author string) ([]*gen.T, error) //返回数据切片和error
}

// Filter 自定义Filter接口
type Filter interface {
	//	select * from @@table where @@column=@value
	FilterWithColumn(column string, value string) ([]*gen.T, error)
}

// Searcher 自定义接口
type Searcher interface {
	// Search 根据指定条件查询书籍

	// SELECT * FROM @@table
	// WHERE publish_date is not null
	// {{if book != nil}}
	//   {{if book.ID > 0}}
	//     AND id = @book.ID
	//   {{else if book.Author != ""}}
	//     AND author=@book.Author
	//   {{end}}
	// {{end}}
	Search(book *gen.T) ([]*gen.T, error)
}

// SearchCustomer 自定义接口
type SearchCustomer interface {
	//	search 条件查询,注意if语法中的book.Author为表对应模型中的字段名，不能直接填写表中的字段名

	//	SELECT
	//	*
	//	FROM
	//	@@table
	//	WHERE
	//	publish_date IS NOT NULL
	// {{if book.Author == ""}}
	//   AND create_at <= @book.CreateAt
	// {{end}}
	SearchCustom(book *gen.T) ([]*gen.T, error)
}
