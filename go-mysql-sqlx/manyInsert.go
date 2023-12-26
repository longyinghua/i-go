package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

// 批量插入方法一
// 自己拼接语句实现批量查询
// BatchInsertUsers 自行构造批量插入的语句
//
// BatchInsertUsers1
//
//	@Description: 自行拼接sql插入语句
//	@param users []*User
//	@return error error
func BatchInsertUsers1(users []*User) error {
	// 存放 (?, ?) 的slice， 用于存放占位符 (?, ?)
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice， 用于存放实际的插入值
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}

	fmt.Printf("valueStrings:%v\n", valueStrings) //  打印出valueStrings [(?,?) (?,?)] sql当中的占位符
	fmt.Printf("valueArgs:%v\n", valueArgs)       //  打印出valueArgs [111 11 222 22]  sql占位符的具体值，按照顺序

	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO sql_test (name, age) VALUES %s",
		strings.Join(valueStrings, ",")) //  %s 用于将slice中的元素拼接成字符串

	fmt.Printf("stmt:%s\n", stmt) //  但因拼接后的sql INSERT INTO sql_test (name, age) VALUES (?, ?),(?, ?)

	_, err := db.Exec(stmt, valueArgs...) //  valueArgs 中的值将按顺序替换占位符中的问号
	fmt.Printf("err:%v\n", err)           //要在控制台上看到错误消息，您可以在调用 db.Exec() 后添加检查 即if err!= nil { fmt.Printf("err:%v\n", err) }
	return err                            //  如果return err 不会再控制台上显示错误消息，而BatchInsertUsers函数调用方接收err后可以printf来处理err显示

	//返回错误：
	//
	//当从函数中返回错误时，它不会自动显示错误或终止程序。
	//调用代码（调用 BatchInsertUsers 的函数）有责任适当地处理该错误。
	//如果调用代码不明确检查错误并处理它，错误将不会显示。
}

// 批量插入方法二
// 使用sqlx.In实现批量插入
// 前提是需要我们的结构体实现driver.Valuer接口

// @Description: 批量插入
// @receiver u User
// @return driver.Value 返回值
// @return error 批量插入失败返回错误
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// 使用sqlx.In实现批量插入代码如下：
// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
//
//	@Description: 批量插入
//	@param users []*User，需要插入的用户切片
//	@return error 批量插入失败返回错误
func BatchInsertUsers2(users []interface{}) error {
	sqlStr := "insert into sql_test (name, age) VALUES (?), (?), (?)"
	query, args, err := sqlx.In(
		sqlStr,   //  这里的（?）, (?) , (?) 占位符表示插入3条数据，与传入参数的个数一致，如果不一致则会报错
		users..., //  如果arg实现了driver。Valuer，sqlx.in会自动调用arg.Value()方法获取值   ,在这里我们自定实现le参数User的Value方法，返回driver.Value类型的值
	)

	fmt.Printf("users:%v\n", users)

	if err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}

	fmt.Println(query) //  查看生成的sql语句
	fmt.Println(args)  //  查看生成的sql语句中的参数

	_, err = db.Exec(query, args...)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}

	return err
}

// 批量插入方法三
// 使用NamedExec实现批量插入
//
// BatchInsertUsers3
//
//	@Description: 批量插入
//	@param users []*User，需要插入的用户切片
//	@return error 批量插入失败返回错误
func BatchInsertUsers3(users []*User) error {
	sqlStr := "insert into sql_test (name, age) VALUES (:name, :age)"
	_, err := db.NamedExec(sqlStr, users)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}

	return err
}

// in 查询
// 查询id在给定id集合中的数据
// SELECT * FROM sql_test WHERE id in (3, 2, 1);
// QueryByIDs 根据给定ID查询
//
//	@Description: 根据给定ID查询
//	@param ids []int，需要查询的id切片
//	@return users []User，查询结果切片
//	@return err error，查询失败返回错误
func QueryByIDs(ids []int) (users []User, err error) {
	//	动态填充id
	sqlStr := "select name,age from sql_test where id IN (?)"
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return nil, err
	}

	//fmt.Printf("query:%s\n", query)
	//fmt.Printf("args:%v\n", args)

	//sqlx.In 返回 `?` bindvar 的查询语句，我们使用Rebind()重新绑定查询语句
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}

	return
}

// in查询和FIND_IN_SET函数
// 查询id在给定id集合的数据并维持给定id集合的顺序
// QueryAndOrderByIDs 按照指定id查询并维护顺序
// QueryAndOrderByIDs
//
//	@Description: in查询，按照指定id排序
//	@param ids []int，需要查询的id切片
//	@return users []User，查询结果切片
//	@return err error，查询失败返回错误
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	//	动态填充id
	strIDs := make([]string, 0, len(ids))

	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}

	sqlStr := "select id,name,age from sql_test where  id IN (?) order by find_in_set(id,?)"

	orderID := strings.Join(strIDs, ",")

	//fmt.Printf("orderID:%s\n", orderID)

	query, args, err := sqlx.In(sqlStr, ids, orderID) //  返回一个新的查询字符串和参数列表
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return nil, err
	}

	//fmt.Printf("query:%s\n", query)
	//fmt.Printf("args:%v\n", args)

	//query 是带有 ? 占位符的查询语句，我们使用 db.Rebind() 重新绑定它，以适应后端数据库

	// sqlx.In 返回带 `?` 如果我们使用 PostgreSQL，占位符应该是 $1，而如果使用 MySQL，占位符应该是 ?
	//db.Rebind() 是一个用于处理不同数据库驱动程序的占位符格式的实用函数，确保我们的查询语句在不同数据库之间保持一致。即重新绑定后mysql占位符为?,postgresql占位符为$1
	query = db.Rebind(query)

	//fmt.Printf("ReBindQuery:%s\n", query)

	err = db.Select(&users, query, args...)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return nil, err
	}

	return
}
