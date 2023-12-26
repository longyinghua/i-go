package main

import (
	"database/sql" //  操作数据库的包
	"fmt"
	_ "github.com/go-sql-driver/mysql" //  数据库类型为mysql的连接驱动，不可缺少，操作不同类型的数据库需要导入不同类型数据库驱动
	"log"
)

// 定义一个全局数据库连接对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "root:dfzy_12345@tcp(192.168.2.250:3306)/test?charset=utf8mb4&parseTime=True"
	// 不会教研账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败：", err)
		return err
	}

	//尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

type User struct {
	id   int
	name string
	age  int
}

// QueryUser
//
//	@Description: SQL单条查询示例
//
// 行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）
func QueryUser() {
	sqlStr := "select * from sql_test where id = ?"
	var u User
	row := db.QueryRow(sqlStr, 1)
	err := row.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d", u.id, u.name, u.age)
}

// QueryUserMany
//
//	@Description: SQL多条查询示例
func QueryUserMany() {
	sqlStr := "select * from sql_test where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	//非常重要，关闭rows对象，释放数据库连接
	defer rows.Close()

	//	循环读取结果集中的数据
	for rows.Next() {
		var u User
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}
}

// InsertUser
//
//	@Description: SQL插入数据示例
func InsertUser() {
	sqlStr := "insert into test.sql_test(name, age) VALUE (?,?)"
	result, err := db.Exec(sqlStr, "liwenzhou", 20)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theId, err := result.LastInsertId() //  获取插入数据的主键id
	if err != nil {
		fmt.Printf("get last insert id failed,err:%v\n", err)
		return
	}
	fmt.Printf("last insert id:%d\n", theId)
}

// UpdateUser
//
//	@Description: SQL更新数据示例
func UpdateUser() {
	sqlStr := "update test.sql_test set age=? where id=?"
	result, err := db.Exec(sqlStr, 21, 4)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rows affected failed,err:%v\n", err)
		return
	}
	fmt.Printf("rows affected:%d\n", rowsAffected)
}

// DeleteUser
//
//	@Description: SQL删除数据示例
func DeleteUser() {
	sqlStr := "delete from test.sql_test where id > ?"
	result, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rows affected failed,err:%v\n", err)
		return
	}
	fmt.Printf("rows affected:%d\n", rowsAffected)

}

//普通SQL语句执行过程：
//
//客户端对SQL语句进行占位符替换得到完整的SQL语句。
//客户端发送完整SQL语句到MySQL服务端
//MySQL服务端执行完整的SQL语句并将结果返回给客户端。
//预处理执行过程：
//
//把SQL语句分成两部分，命令部分与数据部分。
//先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
//然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
//MySQL服务端执行完整的SQL语句并将结果返回给客户端。

//为什么要预处理？
//优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
//避免SQL注入问题。

// PrepareUser
//
//	@Description: go实现的SQL预处理 select
func PrepareUser() {
	sqlStr := "select id, name, age from sql_test where id > ?"
	stmt, err := db.Prepare(sqlStr) //  返回一个准备好的sql语句对象，可以重复使用
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0) //  在预处理sql语句对象增加执行参数，执行sql语句，返回一个结果集对象rows
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	//循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)

	}
}

// 插入，更新，删除数据操作的预处理十分类似，这里以插入操作的预处理为例
//
// PrePareInsertUser
//
//	@Description: go实现的SQL预处理,insert,delete,update用法相同
func PrePareInsertUser() {
	dbStr := "insert into sql_test(name, age) values (?,?)"
	stmt, err := db.Prepare(dbStr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("hahahaha", 20)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	_, err = stmt.Exec("alalalala", 30)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	fmt.Println("insert success")
}

// 我们任何时候都不应该自己拼接SQL语句！
// select id, name, age from sql_test where name='xxx' or 1=1#'      #号表示注释，所有的#开头及后面的都是注释
// sql注入示例
//
// sqlInjectDemo
//
//	@Description: SQL注入示例
//	@param name : SQL占位符中的参数
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from sql_test where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u User
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("init db failed,err:%v\n", err)
		return
	}
	//QueryUser() //
	//InsertUser()
	//UpdateUser()
	//DeleteUser()
	//PrepareUser()
	//PrePareInsertUser()
	QueryUserMany()
	//sqlInjectDemo("xxx' or 1=1#")
	transation() //  事务调教demo
}
