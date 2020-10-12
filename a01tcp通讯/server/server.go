package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端异常退出", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {

}
