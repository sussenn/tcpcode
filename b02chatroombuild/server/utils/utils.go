package utils

import (
	message "mygithub/tcpcode/b02chatroombuild/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (transfer *Transfer) ReadPkg(mes message.Message, err error) {

}
