package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

//redis连接池的初始化
func initPool(maxIdle, maxActive int, idleTimeout time.Duration, address string) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
