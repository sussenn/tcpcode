package main

import (
	"fmt"
	"os"
	"tcpcode/b02chatroombuild/client/processCli"
)

//账户 密码
var userId int
var userPwd string
var userName string

func main() {
	//接收 键盘输入的选择
	var key int
	for {
		fmt.Println("--------------欢迎登录--------------")
		fmt.Println("\t\t\t 1 登录")
		fmt.Println("\t\t\t 2 注册")
		fmt.Println("\t\t\t 3 退出")
		fmt.Println("请选择(1-3):")
		//接收键盘输入
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("[登录聊天室]")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			//登录
			up := &processCli.UserProcess{}
			_ = up.Login(userId, userPwd)
		case 2:
			fmt.Println("[注册用户]")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户昵称")
			fmt.Scanf("%s\n", &userName)
			//注册
			up := &processCli.UserProcess{}
			_ = up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误,请重新选择")
		}
	}
}
