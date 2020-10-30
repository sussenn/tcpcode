package message

//最终响应 消息体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//登录响应结果 消息体
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	//保存在线用户id的切片
	UsersId []int
}

//登录请求 消息体
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//注册 消息体
type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//通知用户在线 消息体 -服务器推送的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//用户状态	iota:从0开始,往下递增. 即UserOnline = 0, UserOffline = 1 .....
const (
	//上线
	UserOnline = iota
	//离开
	UserOffline
	//繁忙
)

//短消息
type SmsMes struct {
	//消息内容
	Content string `json:"content"`
	User
}
