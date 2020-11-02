package main

import (
	"fmt"
	//"mygithub/tcpcode/b02chatroombuild/server/model"
	"net"
	"tcpcode/b02chatroombuild/server/model"
	"time"
)

//实例化userDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//初始化redis
	initPool(8, 0, 40*time.Second, "localhost:6379")
	//实例化userDao
	initUserDao()
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
	//调用总控
	process := &Processor{
		Conn: conn,
	}
	err := process.processTwo()
	if err != nil {
		fmt.Println("客户端和服务端连接协程异常. err: ", err)
		return
	}
}
