package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"mygithub/tcpcode/b01chatroom/common/message"
	"net"
)

func main() {
	//启动服务器
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器启动监听失败. err: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器监听连接失败. err: ", err)
		}
		//服务器启动成功, 开启协程进行客户端连接通讯
		go process(conn)
	}
}

//客户端连接
func process(conn net.Conn) {
	defer conn.Close()
	for {
		//读取客户端消息 的函数
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端连接服务器关闭...")
				return
			} else {
				fmt.Println("读取客户端消息失败. err:", err)
				return
			}
		}
		//根据客户端发送消息类型不同, 调用相应函数
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

//读客户端消息
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("客户端发送消息...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//自定义异常: 读包, header头出错
		fmt.Println("读包, header出错. err:", err)
		return
	}
	//先获取客户端发送数据的长度,避免发生丢包
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据发送数据的长度,读取消息
	//代码解释: 读取pkgLen长度的数据放入到buf切片
	n, err := conn.Read(buf[:pkgLen])
	//n是int类型
	if n != int(pkgLen) || err != nil {
		//自定义异常: 读包, body出错
		fmt.Println("读包, body出错. err:", err)
		return
	}
	//将buf反序列化成mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("readPkg()反序列化失败. err: ", err)
		return
	}
	return
}

//根据客户端发送消息类型不同, 调用相应函数
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//登录处理
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	//注册处理
	default:
		fmt.Println("消息体类型不匹配...")
	}
	return err
}

//登录处理
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//声明登录请求消息体
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("serverProcessLogin() '登录请求消息'反序列化失败. err=", err)
		return
	}
	//声明最终响应消息体的类型 ->登录响应结果消息体类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//声明登录响应消息体
	var loginResMes message.LoginResMes

	//登录验证
	if loginMes.UserId == 100 && loginMes.UserPwd == "123" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "账户名密码错误"
	}

	//先将登录响应结果序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("'登录响应消息体'序列化失败. err=", err)
		return
	}
	//将序列化后的'登录响应结果'赋值给'最终响应消息体'
	resMes.Data = string(data)
	//将'最终响应消息体'序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("'最终响应消息体'序列化失败. err=", err)
		return
	}
	//发送响应
	err = writePkg(conn, data)
	return
}

//发送响应
func writePkg(conn net.Conn, data []byte) (err error) {
	//发送长度,验证连接
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	fmt.Println("服务端发送长度内容验证... buf: ", buf)
	//发送数据长度信息
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("连接失败. err: ", err)
		return
	}
	//发送消息
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送消息失败. err:", err)
		return
	}
	return
}
