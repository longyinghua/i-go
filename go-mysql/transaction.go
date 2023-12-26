package main

import "fmt"

// transation
//
//	@Description: SQL事务提交demo
func transation() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin transaction error: %v\n", err)
		return
	}

	//第一条sql
	sqlStr1 := "update sql_test set age=88 where id = ?"
	result, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		tx.Rollback() //  执行update sql失败，回滚事务
		fmt.Printf("insert sql error: %v\n", err)
		return
	}
	rowsAffected1, err := result.RowsAffected()
	if err != nil {
		tx.Rollback() //  但凡任何错误，都回滚事务
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "update sql_test set age=99 where id=?"
	exec, err := tx.Exec(sqlStr2, 2)
	if err != nil {
		tx.Rollback() //  执行update sql失败，回滚事务
		fmt.Printf("insert sql error: %v\n", err)
		return
	}
	rowsAffected2, err := exec.RowsAffected()
	if err != nil {
		tx.Rollback() //  但凡任何错误，都回滚事务
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	//打印影响了多少行
	//fmt.Println(rowsAffected1, rowsAffected2)
	fmt.Printf("sqlStr1执行了%d行\n", rowsAffected1)
	fmt.Printf("sqlStr2执行了%d行\n", rowsAffected2)

	if rowsAffected1 == 1 && rowsAffected2 == 1 {
		fmt.Println("exec trans success! 事务执行成功啦...")
		tx.Commit() //  事务提交
	} else {
		fmt.Println("update failed --- 事务回滚啦...")
		tx.Rollback() //  事务回滚
	}

}
