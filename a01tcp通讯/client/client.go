package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//2. 再启动客户端
//步骤:
//a. 客户端进行拨号连接服务端net.Dial()
//b. 创建输入流bufio.NewReader(os.Stdin),然后输出到服务端conn.Write()
func main() {
	conn, err := net.Dial("tcp", "10.117.0.124:8888")
	if err != nil {
		fmt.Println("客户端拨号失败:", err)
		return
	}
	//客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin: 标准输入[终端]
	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取客户端输入数据失败:", err)
		}
		//如果用户输入的是 exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		//再将line一行输入数据 发送给 服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("客户端数据写入服务端失败:", err)
		}
	}
}
