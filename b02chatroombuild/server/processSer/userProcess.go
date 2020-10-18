package processSer

import (
	"encoding/json"
	"fmt"
	"net"
	message "tcpcode/b02chatroombuild/common"
	"tcpcode/b02chatroombuild/server/model"
	"tcpcode/b02chatroombuild/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

//登录处理
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//声明登录请求消息体
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("'登录请求消息'反序列化失败. err=", err)
		return
	}
	//声明最终响应消息体的类型 ->登录响应结果消息体类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//声明登录响应消息体
	var loginResMes message.LoginResMes

	//登录验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			//err.Error(): 读取错误描述
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		fmt.Println("用户登录成功,用户信息: ", user)
	}
	//先将登录响应结果序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("'登录响应消息体'序列化失败. err=", err)
		return
	}
	//将序列化后的'登录响应结果'赋值给'最终响应消息体'
	resMes.Data = string(data)
	//将'最终响应消息体'序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("'最终响应消息体'序列化失败. err=", err)
		return
	}
	//发送响应
	//创建Transfer实例,调用utils工具
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//注册处理
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("'注册请求消息'反序列化失败. err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	//调用dao 注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("'注册响应消息体'序列化失败. err=", err)
		return
	}
	//将序列化后的'注册响应结果'赋值给'最终响应消息体'
	resMes.Data = string(data)
	//将'最终响应消息体'序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("'最终响应消息体'序列化失败. err=", err)
		return
	}
	//发送响应
	//创建Transfer实例,调用utils工具
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
