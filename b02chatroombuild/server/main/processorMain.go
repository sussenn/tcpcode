package main

import (
	"fmt"
	"io"
	//"mygithub/tcpcode/b02chatroombuild/common"
	//"mygithub/tcpcode/b02chatroombuild/server/processSer"
	//"mygithub/tcpcode/b02chatroombuild/server/utils"
	"net"
	"tcpcode/b02chatroombuild/common"
	"tcpcode/b02chatroombuild/server/processSer"
	"tcpcode/b02chatroombuild/server/utils"
)

//总控
type Processor struct {
	Conn net.Conn
}

//根据客户端发送消息类型不同, 调用相应函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//登录处理
		//创建UserProcess实例,调用方法
		up := &processSer.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//注册处理
		up := &processSer.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//群发消息
		smsProcess := &processSer.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息体类型不匹配...")
	}
	return err
}

//客户端连接服务器
func (this *Processor) processTwo() (err error) {
	for {
		//读取客户端消息 的函数
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端连接服务器关闭...")
				return err
			} else {
				fmt.Println("读取客户端消息失败. err:", err)
				return err
			}
		}
		//根据客户端发送消息类型不同, 调用相应函数
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
