package utils

import (
	"database/sql"
	//数据库驱动,不使用也必须导入
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	//Open函数不创建与数据库的连接,只是检测当前连接参数是否合法
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_stu")
	if err != nil {
		panic(err.Error())
	}
}
