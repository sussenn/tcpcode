package processSer

import (
	"encoding/json"
	"fmt"
	message "mygithub/tcpcode/b02chatroombuild/common"
	"mygithub/tcpcode/b02chatroombuild/server/utils"
	"net"
)

type SmsProcess struct {
}

//转发群聊消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	//取出data
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes() 转发群聊消息 反序列化失败. err: ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes() 转发群聊消息 序列化失败. err: ", err)
		return
	}
	//取出所有在线用户
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		//给在线用户转发消息
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

//转发消息
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser() 转发消息失败. err: ", err)
	}
}
