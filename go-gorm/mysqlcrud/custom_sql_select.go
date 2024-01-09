package mysqlcrud

import (
	"context"
	"go-gorm/dal/model"
	"go-gorm/dal/query"
	"log"
)

// CustomSQLSelect1
//
//	@Description: 自定义SQL查询,根据author名字查询结果
//	@param author string，作者名字
//	@return []*model.Book 返回结果切片
//	@return error 错误信息
func CustomSQLSelect1(author string) ([]*model.Book, error) {
	rets, err := query.Book.WithContext(context.Background()).GetBooksByAuthor(author)
	if err != nil {
		log.Printf("GetBooksByAuthor fail, err:%v\n", err)
		return nil, err
	}

	return rets, nil
}

func CustomSQLSelect2(id int64) (*model.Book, error) {
	ret, err := query.Book.WithContext(context.Background()).GetByID(id)
	if err != nil {
		log.Printf("GetByID fail, err:%v\n", err)
		return nil, err
	}

	return &ret, nil
}

func CustomSQLSelect3(id int64) (map[string]interface{}, error) {
	resultMap, err := query.Book.WithContext(context.Background()).GetByIDReturnMap(id)
	if err != nil {
		log.Printf("GetByIDReturnMap fail, err:%v\n", err)
		return nil, err
	}

	return resultMap, nil
}
