package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

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
	//批量写入Hash类型数据	k-v(map)
	_, err = conn.Do("HMSet", "user001", "name", "Mark", "age", 18)
	if err != nil {
		fmt.Println("conn.Do() HMSet-> err:", err)
		return
	}

	//批量读取Hash类型数据 redis.Strings()读取多个
	r, err := redis.Strings(conn.Do("HMGet", "user001", "name", "age"))
	if err != nil {
		fmt.Println("conn.Do() HMGet-> err:", err)
		return
	}
	//r=[mark 18]
	fmt.Printf("hash数据格式: r=%v\n", r)

	//遍历输出
	for i, v := range r {
		fmt.Printf("r[%d] = %s\n", i, v)
	}
}
