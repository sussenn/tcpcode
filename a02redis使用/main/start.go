package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//GoPath下(即C:/GoProject下,src同级目录) cmd执行go get github.com/garyburd/redigo/redis命令下载redis库
func main() {
	//1.连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial() err:", err)
		return
	}
	fmt.Println("连接redis成功...", conn)
	//关闭
	defer conn.Close()
	//2.写入string类型数据	k-v(string)
	//本质就是redis客户端的操作命令使用
	//name-宿舍森女
	_, err = conn.Do("Set", "name", "宿舍森女")
	if err != nil {
		fmt.Println("conn.Do() Set-> err:", err)
		return
	}
	//3.读取string类型数据
	//不能使用断言获取string类型 err-> r.(string)
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do() Get-> err:", err)
		return
	}
	fmt.Println("获取redis数据成功: ", r)
}
