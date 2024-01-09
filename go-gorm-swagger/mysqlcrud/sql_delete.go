package mysqlcrud

import (
	"context"
	"go-gorm/dal/model"
	"go-gorm/dal/query"
	"log"
)

// DeleteBook1
//
//	@Description:	删除单条数据
//	@param			book	*model.Book
//	@return			int64 rowsAffected返回删除的记录数
//	@return			error error错误信息

func DeleteBook1(book *model.Book) (int64, error) {
	u := query.Book
	ctx := context.Background()
	resultInfo, err := u.WithContext(ctx).Where(u.Author.Eq(book.Author), u.ID.Lte(book.ID)).Delete()
	//resultInfo, err := u.WithContext(ctx).Where(u.Author.Eq(book.Author)).Or(u.ID.Lte(book.ID)).Delete()
	// DELETE from book where author = "abc" AND id <= 10；
	if err != nil {
		log.Printf("delete book error: %v", err)
		return 0, err
	}
	rowsAffected := resultInfo.RowsAffected
	return rowsAffected, nil
}
