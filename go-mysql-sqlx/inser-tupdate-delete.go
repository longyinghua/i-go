package main

import (
	"errors"
	"fmt"
)

//插入、更新、删除
//sqlx包中exec方法和原生sql包中Exec方法使用基本一致

// 插入数据
func InsertRowDemo() {
	sqlStr := "insert into sql_test(name,age) values (?,?)"
	result, err := db.Exec(sqlStr, "张三", 31)
	if err != nil {
		fmt.Printf("插入数据失败，err:%v\n", err)
		return
	}
	theID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("获取插入数据的ID失败，err:%v\n", err)
		return
	}
	fmt.Printf("插入数据成功，ID:%d\n", theID)
}

// 更新数据
func UpdateRowDemo() {
	sqlStr := "update sql_test set age=? where id=?"
	result, err := db.Exec(sqlStr, 32, 100)
	if err != nil {
		fmt.Printf("更新数据失败，err:%v\n", err)
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("获取更新数据的影响行数失败，err:%v\n", err)
		return
	}
	fmt.Printf("更新数据成功，影响行数:%d\n", rows)
}

// 删除数据
func DeleteRowDemo() {
	sqlStr := "delete from sql_test where id > ?"
	result, err := db.Exec(sqlStr, 100)
	if err != nil {
		fmt.Printf("删除数据失败，err:%v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("获取删除数据的影响行数失败，err:%v\n", err)
		return
	}
	fmt.Printf("删除数据成功，影响行数:%d\n", rowsAffected)
}

// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段。
func InsertUserDemo() (err error) {
	//sqlx 提供了一种更方便的方式，即使用命名参数。这样，您可以在 SQL 查询中使用具有名称的占位符，例如 :name 和 :age
	sqlStr := "insert into sql_test(name,age) values (:name,:age)" //  DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段

	// user := User{
	//	 Name: "allalalalal",
	//	 Age:  100,
	// }
	//db.NamedExec(sqlStr, user)
	db.NamedExec(sqlStr, User{
		Name: "qimi",
		Age:  20,
	})
	//db.NamedExec(
	//	sqlStr,
	//	map[string]interface{}{
	//		"name": "qimi",
	//		"age":  20,
	//	},
	//)
	return
}

//数据库连接的生命周期：
//
//在执行查询操作时，我们从数据库连接池中获取一个连接。
//对于查询操作，我们需要手动关闭连接，以便将连接返回给连接池，以便其他请求可以继续使用它。
//但是，对于插入、更新或删除等操作，sqlx 会在执行完操作后自动关闭连接。

// NamedQuery与DB.NamedExec同理，这里是支持查询。
func nameQueryDemo() {
	sqlStr := "select * from sql_test where  name=:name"

	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	user := User{
		Age:  30,
		Name: "qimi",
	}
	rows, err := db.NamedQuery(
		sqlStr,
		user,
	)
	if err != nil {
		fmt.Printf("查询数据失败，err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		//err := rows.Scan(&user.Id, &user.Name, &user.Age)
		err := rows.StructScan(&user)
		if err != nil {
			fmt.Printf("转换数据到struct失败，err:%v\n", err)
			return
		}
		fmt.Printf("查询到的数据:%+v\n", user)
	}

	// 使用map做命名查询
	result, err := db.NamedQuery(sqlStr, map[string]interface{}{
		"name": "qimi",
	})
	if err != nil {
		fmt.Printf("查询数据失败，err:%v\n", err)
		return
	}

	defer result.Close()

	for result.Next() {
		err := result.StructScan(&user)
		if err != nil {
			fmt.Printf("转换数据到struct失败，err:%v\n", err)
			return
		}
		fmt.Printf("查询到的数据:%+v\n", user)
	}
}

// 事务提交2
func transactionDemo2() (err error) {
	tx, err := db.Beginx() //  开启事务
	if err != nil {
		fmt.Printf("开启事务失败，err:%v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback() //  回滚事务,但凡有任何错误，都要回滚事务
			panic(p)      //  抛出错误
		} else if err != nil {
			fmt.Println("rollback-----")
			tx.Rollback()
		} else {
			err = tx.Commit() //  在同一个事务中，如果没有错误，则提交事务，这里err后面要使用=号，而不能用:=,因为上面已经定义了变量err
			if err != nil {
				fmt.Println("commit")
			}
		}
	}()

	sqlStr1 := "update sql_test set age=20 where id=?"

	result, err := tx.Exec(sqlStr1, 100)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("exec sqlStr1 failed-------")
	}

	sqlStr2 := "update sql_test set age=50 where id=?"
	rs, err := tx.Exec(sqlStr2, 101)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr2 failed-------")
	}

	return err
}
