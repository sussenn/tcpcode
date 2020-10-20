package processCli

import (
	"encoding/json"
	"fmt"
	"mygithub/tcpcode/b02chatroombuild/client/utils"
	message "mygithub/tcpcode/b02chatroombuild/common"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	//短消息内容
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes() 发送短消息 序列化失败. err: ", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes() 发送短消息 最终请求体序列化失败. err: ", err)
		return
	}
	//将消息发送给服务端
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes() 发送短消息到服务端失败. err: ", err)
		return
	}
	return
}
