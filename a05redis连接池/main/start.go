package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//全局连接池
var pool *redis.Pool

//在主函数运行前初始化
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,  //最大空闲连接数
		MaxActive:   0,  //最大连接数. 0:无限制
		IdleTimeout: 40, //连接最大空闲时间(超时即回收)
		//初始化
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//获取一个连接
	conn := pool.Get()
	defer conn.Close()
	//存
	_, err := conn.Do("Set", "key01", "大将军")
	if err != nil {
		fmt.Println("conn.Do() err: ", err)
		return
	}
	//取
	r, err := redis.String(conn.Do("Get", "key01"))
	if err != nil {
		fmt.Println("conn.Do() err: ", err)
		return
	}
	fmt.Println("redis 取出结果 r = ", r)
}
