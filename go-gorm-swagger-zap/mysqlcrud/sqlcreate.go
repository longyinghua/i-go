package mysqlcrud

import (
	"context"
	"go-gorm-swagger-zap/dal/model"
	"go-gorm-swagger-zap/dal/query"
)

// CreateBook1
//
//	@Description:	插入单条数据
//	@param			book	*model.Book	参数为结构体模型，可以根据请求体中的body绑定模型结构体后取具体的字段作为SQL条件进行insert插入
//	@return			error 返回错误信息

func CreateBook1(book *model.Book) error {

	err := query.Book.WithContext(context.Background()).Create(book) //  插入单条数据

	if err != nil {
		//log.Printf("create book error: %v\n", err)
		return err
	}

	return nil
}

//更高效地插入大量记录，可以将一个 slice 传递给 Create 方法。 将切片数据传递给 Create 方法，GORM 将生成单个 SQL 语句来插入所有数据，并回填主键的值
//var users = []*model.User{{Name: "modi"}, {Name: "zhangqiang"}, {Name: "songyuan"}}
//query.User.WithContext(ctx).Create(users...)

// CreateBook2
//
//	@Description:	批量插入数据
//	@param			books	[]*model.Book	参数为Book列表
//	@return			error 返回错误信息

func CreateBook2(books []*model.Book) error {
	//[]*model.Book形参，*model.Book为接头体指针切片，实参表示为books...,表示多个*model.Book类型组成的切片元素，当不知道有多少个元素时，使用...表示
	u := query.Book
	c := context.Background()

	err := u.WithContext(c).Create(books...) // ...*mode.book 表示*model.Book类型切片,*model.Book结果集
	if err != nil {
		//log.Printf("create book error: %v\n", err)
		return err
	}

	return nil
}

//1、在Go中，... 是一个省略号操作符，被称为“variadic ellipsis”，用于将切片中的元素展开为单独的参数。在这个特定的函数中，Create(books...) 的语法表示将切片 books 中的所有元素作为单独的参数传递给 Create 函数。
//
//2、在这个场景中，Create 函数可能期望接收多个参数，每个参数都是一个 *model.Book 类型的指针。而 books... 的语法确保将 books 中的每个 *model.Book 类型的指针都传递给 Create 函数作为单独的参数。
//
//3、换句话说，Create(books...) 可以理解为将切片 books 中的所有元素展开，作为单独的参数传递给 Create 函数。这样的写法通常用于处理函数参数为可变数量参数的情况。
