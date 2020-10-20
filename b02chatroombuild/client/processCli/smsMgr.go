package processCli

import (
	"encoding/json"
	"fmt"
	message "mygithub/tcpcode/b02chatroombuild/common"
)

//客户端显示群发消息
func outputGroupMes(mes *message.Message) {
	//读取信息
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes() 反序列化失败. err: ", err.Error())
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id: %d 对大家说: %s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
