package mysqlcrud

import (
	"context"
	"fmt"
	"go-gorm/dal/query"
)

func UpdateBook() {
	//	更新表中某条数据
	resultInfo, err := query.Book.WithContext(context.Background()).Where(query.Book.ID.Eq(3)).Update(query.Book.Price, 200) //  where单条件，WHERE `book`.`id` = 3
	//resultInfo, err := query.Book.WithContext(context.Background()).Where(query.Book.ID.Eq(3), query.Book.Price.Eq(300)).Update(query.Book.Price, 200) //  where多条件，WHERE `book`.`id` = 3 AND `book`.`price` = 300

	if err != nil {
		fmt.Printf("update book error: %v\n", err)
		return
	}
	fmt.Printf("RowsAffected book result: %v\n", resultInfo.RowsAffected)

}
