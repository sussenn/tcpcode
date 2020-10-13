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
	//写入Hash类型数据	k-v(map)
	//user01- name:tom  	 user02- age:18
	_, err = conn.Do("HSet", "user01", "name", "tom")
	if err != nil {
		fmt.Println("conn.Do() HSet-> err:", err)
		return
	}
	_, err = conn.Do("HSet", "user02", "age", 18)
	if err != nil {
		fmt.Println("conn.Do() HSet-> err:", err)
		return
	}

	//读取Hash类型数据
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("conn.Do() HGet-> err:", err)
		return
	}
	//int类型...
	r2, err := redis.Int(conn.Do("HGet", "user02", "age"))
	if err != nil {
		fmt.Println("conn.Do() HGet-> err:", err)
		return
	}
	fmt.Printf("获取redis Hash数据成功: r1=%v\t r2=%v\t", r1, r2)
}
