package processCli

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"tcpcode/b02chatroombuild/client/utils"
	message "tcpcode/b02chatroombuild/common"
)

type UserProcess struct {
}

//登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//1.连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接服务器失败. err: ", err)
		return
	}
	defer conn.Close()
	//2.定义传输消息体
	var mes message.Message
	//消息体的类型-> 登录信息类型
	mes.Type = message.LoginMesType
	//消息体的数据-> 登录信息
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//3.序列化 loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("login() 序列化失败. err: ", err)
		return
	}
	//将data强转为string,并赋给mes	[空接口强转使用断言]
	mes.Data = string(data)

	//4.将mes对象序列化	-json.Marshal()返回[]byte, 即data是[]byte类型
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("login() 序列化失败. err: ", err)
		return
	}
	//5.发送登录请求
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Login() 客户端发送登录请求消息失败. err: ", err)
		return
	}
	//6.处理响应结果消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Login() 客户端读取响应结果失败. err=", err)
		return
	}
	//7.将响应结果的data反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//开启协程和服务器保持通讯,用于接收服务器推送的消息
		go serverProcessMes(conn)
		//fmt.Println("登录成功")
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

//注册
func (this *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	//1.连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接服务器失败. err: ", err)
		return
	}
	defer conn.Close()
	//2.定义传输消息体
	var mes message.Message
	//消息体的类型-> 注册信息类型
	mes.Type = message.RegisterMesType
	//消息体的封装数据-> 注册信息
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//3.序列化 registerMes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("Register() 注册消息体 序列化失败. err: ", err)
		return
	}
	//将data强转为string,并赋给mes	[空接口强转使用断言]
	mes.Data = string(data)

	//4.将mes对象序列化 -json.Marshal()返回[]byte, 即data是[]byte类型
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Register() 最终请求消息体 序列化失败. err: ", err)
		return
	}

	//5.发送注册请求
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Register() 客户端发送注册请求消息失败. err: ", err)
		return
	}
	//6.接受注册响应
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Register() 客户端读取响应结果失败. err=", err)
		return
	}
	//7.将响应结果的data反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		//注册成功
		fmt.Println("注册成功,请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
