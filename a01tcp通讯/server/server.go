package main

import (
	"fmt"
	"net"
)

//1. 先启动服务端
//步骤:
//a. 根据协议,ip端口号 创建监听器net.Listen()
//b. 监听器接收客户端连接listen.Accept()
//c. 服务端和客户端连接后,开始读取信息conn.Read()
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息. 如果客户端没有write[发送],那么协程就阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端异常退出", err)
			return
		}
		//显示客户端发送到本服务端的内容
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务端监听...")
	//0.0.0.0:8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("服务端监听异常 err:", err)
		return
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务端连接异常 err:", err)
		} else {
			fmt.Printf("服务端连接客户端成功, conn:%v 客户端ip:%v\n", &conn, conn.RemoteAddr().String())
		}
		//接收客户端数据
		go process(conn)
	}
}
