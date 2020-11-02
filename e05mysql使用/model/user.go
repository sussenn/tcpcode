package model

import (
	"fmt"
	"tcpcode/e05mysql使用/utils"
)

type User struct {
	ID       int
	Username string
	Sex      string
	Password string
	Email    string
}

//新增方法(带预编译)
func (user *User) AddUser() error {
	sqlStr := "INSERT INTO `go_stu`.`user` (`name`, `sex`, `password`, `email`) VALUES (?, ?, ?, ?)"
	//预编译 (即先处理sql)
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常: ", err)
		return err
	}
	//执行sql
	_, sqlErr := inStmt.Exec(user.Username, user.Sex, user.Password, user.Email)
	if sqlErr != nil {
		fmt.Println("执行sql异常: ", err)
		return err
	}
	return nil
}

//新增方法(不带预编译)
func (user *User) InsertUser() error {
	sqlStr := "INSERT INTO `go_stu`.`user` (`name`, `sex`, `password`, `email`) VALUES (?, ?, ?, ?)"
	_, sqlErr := utils.Db.Exec(sqlStr, user.Username, user.Sex, user.Password, user.Email)
	if sqlErr != nil {
		fmt.Println("执行sql异常: ", sqlErr)
		return sqlErr
	}
	return nil
}

//单条结果的查询
func (user *User) FindById() (*User, error) {
	sqlStr := "SELECT * FROM `go_stu`.`user` WHERE `id` = ?"
	//QueryRow()查询单条结果,返回值需二次封装
	row := utils.Db.QueryRow(sqlStr, user.ID)
	//将查询的结果封装
	var id int
	var username string
	var sex string
	var password string
	var email string
	err := row.Scan(&id, &username, &sex, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Sex:      sex,
		Password: password,
		Email:    email,
	}
	return u, nil
}

//查询多条结果	切片相当于list
func (user *User) GetUserList() ([]*User, error) {
	sqlStr := "SELECT * FROM `go_stu`.`user`"
	//获得多条记录
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//遍历封装
	var userList []*User
	for rows.Next() {
		var id int
		var username string
		var sex string
		var password string
		var email string
		err := rows.Scan(&id, &username, &sex, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Sex:      sex,
			Password: password,
			Email:    email,
		}
		userList = append(userList, u)
	}
	return userList, nil
}
