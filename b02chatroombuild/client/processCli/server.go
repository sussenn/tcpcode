package processCli

import (
	"encoding/json"
	"fmt"
	"mygithub/tcpcode/b02chatroombuild/client/utils"
	message "mygithub/tcpcode/b02chatroombuild/common"
	"net"
	"os"
	//"tcpcode/b02chatroombuild/client/utils"
)

func ShowMenu(userId int) {
	//fmt.Println("\t\t\t恭喜xxx登录成功")
	fmt.Printf("\t\t用户%d 登录成功\n", userId)
	fmt.Println("\t\t1. 显示在线用户列表")
	fmt.Println("\t\t2. 发送消息")
	fmt.Println("\t\t3. 信息列表")
	fmt.Println("\t\t4. 退出系统")
	fmt.Println("请选择(1-4):")
	var key int
	var content string

	sp := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//显示在线用户列表
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说的什么:")
		fmt.Scanf("%s\n", &content)
		_ = sp.SendGroupMes(content)
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
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//有人上线
			var notifyUserStatusMes message.NotifyUserStatusMes
			_ = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//将当前用户信息状态 保存到map. (用户上线后,更新在线用户map的数据)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//有人群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
		//fmt.Println(mes)
	}
}
