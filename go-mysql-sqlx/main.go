package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //  数据库类型为mysql的连接驱动，不可缺少，操作不同类型的数据库需要导入不同类型数据库驱动
	"github.com/jmoiron/sqlx"
)

// 连接数据库
var db *sqlx.DB

// initDB
//
//	@Description: 初始化数据库连接
//	@return err 连接失败返回错误
func initDB() (err error) {
	dsn := "root:dfzy_12345@tcp(192.168.2.250:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	dbc, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db error: %v\n", err)
		return
	}
	db = dbc
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)

	return
}

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

//导出字段：
//在 Go 中，以大写字母开头的字段被视为导出的（即可在包外访问）。
//使用 sqlx 时，它期望目标结构体中有导出的字段，所以在使用sqlx时，将执行sql后把sql结果集中的字段赋值给结构体对应的字段，所以这里定义的结构体必须是大写

func QueryRowDemo() {
	sqlStr := "select id,name,age from sql_test where id = ?"
	var u User
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("query row error: %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age) //  取变量的值
	//fmt.Printf("id:%d name:%s age:%d\n", &u.Id, &u.Name, &u.Age) //  取变量的地址
}

func QueryMultiRowDemo() {
	sqlStr := "select id,name,age from sql_test where id > ?"
	var users []User
	err := db.Select(&users, sqlStr, 0) //  执行sql查询语句sqlStr，查询参数为0，将sql查询结构赋值给users切片
	if err != nil {
		fmt.Printf("query row error: %v\n", err)
		return
	}

	//fmt.Printf("users:%v\n", users)
	for _, user := range users {
		fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
	}

}

// 更新数据
func UpdateRowDemo1() {
	sqlStr := "update sql_test set age=? where id=?"
	result, err := db.Exec(sqlStr, 33, 101)
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

func SelectR1(id int) {
	dbStr := "select * from sql_test where id = ?"
	query, err := db.Query(dbStr, id)
	if err != nil {
		fmt.Printf("查询数据失败，err:%v\n", err)
		return
	}
	var user User
	for query.Next() {
		err := query.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("查询数据失败，err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
	}
}

func main() {
	initDB()
	//QueryRowDemo()
	fmt.Println("==========================")
	//SelectR1(101)
	//QueryMultiRowDemo()
	//UpdateRowDemo1()
	//nameQueryDemo()

	//BatchInsertUsers1([]*User{&User{Id: 100, Name: "111", Age: 11}, &User{Id: 101, Name: "222", Age: 22}})
	//BatchInsertUsers2([]interface{}{&User{Name: "111", Age: 11}, &User{Name: "222", Age: 22}, &User{Name: "333", Age: 66}})    //  []interface中传入了三个user结构体,也就是三个元素
	//BatchInsertUsers3([]*User{&User{Name: "111", Age: 11}, &User{Name: "222", Age: 22}, &User{Name: "333", Age: 66}})

	//u1 := User{Name: "七米", Age: 18}
	//u2 := User{Name: "q1mi", Age: 28}
	//u3 := User{Name: "小王子", Age: 38}

	//// 方法1
	//users := []*User{&u1, &u2, &u3}
	//err := BatchInsertUsers1(users)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	//}

	//方法2
	//users2 := []interface{}{u1, u2, u3}
	//err = BatchInsertUsers2(users2)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	//}

	//// 方法3
	//users3 := []*User{&u1, &u2, &u3}
	//err = BatchInsertUsers3(users3)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	//}

	fmt.Println("==========================") //  in查询按照指定id顺序排序
	results, err := QueryAndOrderByIDs([]int{101, 102, 103})
	if err != nil {
		fmt.Printf("QueryAndOrderByIDs failed, err:%v\n", err)
		return
	}

	fmt.Printf("res:%v\n", results)

	for _, user := range results {
		fmt.Printf("id:%d name:%s age:%d\n", user.Id, user.Name, user.Age)
	}

}
