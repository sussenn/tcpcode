package model

import (
	message "mygithub/tcpcode/b02chatroombuild/common"
	"net"
)

//存储用户tcp通讯连接 和 用户信息
type CurUser struct {
	Conn net.Conn
	message.User
}
