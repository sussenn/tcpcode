package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	//"tcpcode/b01chatroom/common/message"
	"mygithub/tcpcode/b01chatroom/common/message"
	"net"
)

//读客户端消息 的函数
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("客户端发送消息...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//自定义异常: 读包, header头出错
		//err = errors.New("readPkg() header err")
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
		//err = errors.New("readPkg() body err")
		return
	}
	//将buf反序列化成mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("buf反序列化失败. err: ", err)
	}
	return
}

//发送消息
func writePkg(conn net.Conn, data []byte) (err error) {
	//发送长度,验证连接
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
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
