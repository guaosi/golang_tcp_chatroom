package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, timeout time.Duration) {
	//常用配置
	pool = &redis.Pool{
		MaxIdle:     maxIdle,   //最大空闲链接数
		MaxActive:   maxActive, // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: timeout,   // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个 ip 的 redis
			return redis.Dial("tcp", address)
		},
	}
}
