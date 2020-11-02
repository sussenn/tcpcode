package model

import (
	"fmt"
	"testing"
)

//model下执行命令:go test		go test -v 查看详细的测试过程
func TestMain(m *testing.M) {
	fmt.Println("调用测试函数前执行操作...")
	m.Run()
}

//函数名首字母T大写作为主测试测试,小写即作为子测试程序
func TestUser(t *testing.T) {
	fmt.Println("验证调用子测试程序")
	t.Run("测试子程序test_AddUser", test_AddUser)
}

func test_AddUser(t *testing.T) {
	//user := &User{
	//	Username: "sussenn",
	//	Sex:      "1",
	//	Password: "123",
	//	Email:    "sussenn@qq.com",
	//}
	//_ = user.AddUser()

	user2 := &User{
		Username: "anny01",
		Sex:      "0",
		Password: "bbb",
		Email:    "anny01@qq.com",
	}
	_ = user2.InsertUser()
}

func TestUser_FindById(t *testing.T) {
	user := &User{
		ID: 1,
	}
	u, err := user.FindById()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(u)
}

func TestGetUserList(t *testing.T) {
	user := &User{}
	userList, err := user.GetUserList()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, u := range userList {
		fmt.Println(u)
	}
}
