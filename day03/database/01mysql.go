package main

import (
	// "database/sql"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// var db *sqlx.DB

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	dsn := "root:123456@(127.0.0.1:3380)/ljd"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect failed!")
	}
	// res, err := db.Exec("create table ljd(name varchar(64), age int)")
	// res.RowsAffected()
	err = db.Ping()
	queryMultiRowDemo(db)
	if err != nil {
		fmt.Println("uuuuu")
	}
	defer db.Close()
}

func perpareInsert(db *sql.DB) {
	sqlStr := `insert into ljd(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr) // 预编译处理
	if err != nil {
		fmt.Printf("prepare failed, errL%v\n", err)
		return
	}
	// 只需要传变量
	result, err := stmt.Exec("liujiadong", 29)
	if err != nil {
		fmt.Printf("something is wrong!! %v\n", err)
	}
	fmt.Println(result.RowsAffected())
}

// sqlx 的使用
func queryMultiRowDemo(db *sqlx.DB) {
	sqlStr := "select name, age from ljd"
	var u User
	err := db.Get(&u, sqlStr)
	fmt.Println(u)
	var users []User
	err = db.Select(&users, sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		fmt.Println(user.Name)
	}

}
