package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"tcpcode/b01chatroom/common/message"
	//"mygithub/tcpcode/b01chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
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

	//4.将mes对象序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("login() 序列化失败. err: ", err)
		return
	}
	//5.data即可用于网络传输
	//5.1先将数据长度发送给服务器  conn.Write()需要的是切片类型参数,故得将data长度转化
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	//将消息长度放入buf切片
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//5.2发送数据长度信息
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("客户端发送数据长度验证失败. err: ", err)
		return
	}
	fmt.Printf("客户端发送数据长度: %d, 内容: %s", len(data), string(data))
	//6.发送消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("客户端发送数据失败. err:", err)
		return
	}
	//7.处理响应结果消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("客户端读取响应结果失败. err=", err)
		return
	}
	//7.1将响应结果的data反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
