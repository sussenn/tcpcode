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
	fmt.Println("conn success...", conn)
}
