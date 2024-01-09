package mysqlcrud

import (
	"context"
	"fmt"
	"go-gorm/dal/query"
)

func DeleteBook() {
	//	删除数据表中数据
	ret, err := query.Book.WithContext(context.Background()).Where(query.Book.ID.Eq(4)).Delete()
	if err != nil {
		fmt.Printf("delete book error: %v\n", err)
		return
	}
	fmt.Printf("RowsAffected book result: %v\n", ret.RowsAffected)
}
