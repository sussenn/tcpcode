package main

import (
	"fmt"
	"os"
)

//账户 密码
var userId int
var userPwd string

func main() {
	//接收 键盘输入的选择
	var key int
	//判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("--------------欢迎登录--------------")
		fmt.Println("\t\t\t 1 登录")
		fmt.Println("\t\t\t 2 注册")
		fmt.Println("\t\t\t 3 退出")
		fmt.Println("\t\t\t 请选择(1-3)")
		//接收键盘输入
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误,请重新选择")
		}
	}
	//登录
	if key == 1 {
		fmt.Println("请输入用户ID")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)
		//都在main包,故可以直接调用login.go 函数
		login(userId, userPwd)
	} else if key == 2 {
		fmt.Println("注册...")
	}
}
