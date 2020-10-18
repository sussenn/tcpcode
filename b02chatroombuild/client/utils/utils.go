package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	message "tcpcode/b02chatroombuild/common"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

//读客户端消息
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("客户端发送消息...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//自定义异常: 读包, header头出错
		fmt.Println("读包, header出错. err:", err)
		return
	}
	//先获取客户端发送数据的长度,避免发生丢包
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据发送数据的长度,读取消息
	//代码解释: 读取pkgLen长度的数据放入到buf切片
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	//n是int类型
	if n != int(pkgLen) || err != nil {
		//自定义异常: 读包, body出错
		fmt.Println("读包, body出错. err:", err)
		return
	}
	//将buf反序列化成mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("readPkg()反序列化失败. err: ", err)
		return
	}
	return
}

//发送请求
func (this *Transfer) WritePkg(data []byte) (err error) {
	//发送长度,验证连接
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	fmt.Println("服务端发送长度内容验证... Buf长度: ", len(this.Buf))
	//发送数据长度信息
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("连接失败. err: ", err)
		return
	}
	//发送消息
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送消息失败. err:", err)
		return
	}
	return
}
