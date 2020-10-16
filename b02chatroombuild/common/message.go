package message

//最终响应 消息体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

//登录响应结果 消息体
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//登录请求 消息体
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//注册 消息体
type RegisterMes struct {
	//...
}
