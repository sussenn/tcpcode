package processCli

import (
	"fmt"
	"net"
	"os"
	"tcpcode/b02chatroombuild/client/utils"
)

func ShowMenu() {
	fmt.Println("\t\t\t恭喜xxx登录成功")
	fmt.Println("\t\t\t1. 显示在线用户列表")
	fmt.Println("\t\t\t2. 发送消息")
	fmt.Println("\t\t\t3. 信息列表")
	fmt.Println("\t\t\t4. 退出系统")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	//因为总会使用到SmsProcess实例,因此将其定义在switch外部
	//smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//显示在线用户列表
		//outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说的什么:)")
		fmt.Scanf("%s\n", &content)
		//smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统...")
		os.Exit(0)
	default:
		fmt.Println("输入的选项不正确..")
	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		//客户端正在读取服务器推送的消息...
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMes() 客户端读取服务器推送消息失败. err: ", err)
			return
		}
		fmt.Println(mes)

	}
}
