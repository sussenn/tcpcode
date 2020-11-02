package processCli

import (
	"fmt"
	"tcpcode/b02chatroombuild/client/model"
	message "tcpcode/b02chatroombuild/common"
	//"mygithub/tcpcode/b02chatroombuild/client/model"
	//message "mygithub/tcpcode/b02chatroombuild/common"
)

//客户端维护的map(存放当前在线用户)
var onlineUsers = make(map[int]*message.User, 10)

//登录成功后即初始化CurUser存放用户tcp连接和用户信息
var CurUser model.CurUser

//用户上线后,更新在线用户map的数据
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	//如果 在线用户map中没有该上线用户信息
	if !ok {
		user = &message.User{
			//新增用户
			UserId: notifyUserStatusMes.UserId,
		}
	}
	//如果 在线用户map中原先有该上线用户信息,只更新该用户 在线状态即可
	user.UserStatus = notifyUserStatusMes.Status
	//存入 在线用户map中
	onlineUsers[notifyUserStatusMes.UserId] = user
	//显示当前所有在线用户
	outputOnlineUser()
}

//在客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}
