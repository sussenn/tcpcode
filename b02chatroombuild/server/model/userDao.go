package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	message "mygithub/tcpcode/b02chatroombuild/common"
	//message "tcpcode/b02chatroombuild/common"
)

//全局userDao,方便main.go中实例化
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return
}

//操作redis获取用户信息
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			//用户不存在
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &message.User{}
	//将string类型json转为user对象
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("getUserById 反序列化失败. err: ", err)
		return
	}
	return
}

//登录校验
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	//从redis连接池获取连接
	conn := this.pool.Get()
	defer conn.Close()
	//根据userId获取用户信息
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//注册处理
func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	//redis取到数据,说明用户已存在,禁止注册
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		//用户已存在
		err = ERROR_USER_EXISTS
		return
	}
	//将user对象序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("序列化用户注册信息失败. err: ", err)
		return
	}
	//保存到redis
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用户注册信息失败. err: ", err)
		return
	}
	return
}
