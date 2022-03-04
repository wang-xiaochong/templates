package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

// Setup redis库初始化
func Setup() *redis.Pool {
	redisPool = &redis.Pool{
		 MaxIdle:     3,
		 MaxActive:   8,
		 IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c,err := redis.Dial("tcp","localhost:6379",redis.DialDatabase(1))	//本地无密码连接
			// c,err := redis.Dial("tcp","82.156.109.119:6379",redis.DialDatabase(1),redis.DialPassword("密码"))	
			if err!=nil {
				fmt.Println("连接失败")
				return nil, err
			}
				return c,err
		},
	}
	return redisPool
}

// GetRedisPool 获取库连接
func GetRedisPool() *redis.Pool {
	if redisPool==nil {
		redisPool= Setup()
	}
	return redisPool
}


// SetKey 设置键值
func SetKey(key string, value string) error {
	var Conn = GetRedisPool().Get()
	//defer Conn.Close()
	_,err:=Conn.Do("set",key,value,"ex",24*60*60)
	if err!=nil {
		return err
	}
	return nil
}

// GetKey 获取键值
func GetKey(key string) ( string, error) {
	var Conn = GetRedisPool().Get()
	//defer Conn.Close()
	ret, err := redis.String(Conn.Do("get", key))
	if err!=nil {
		return ret,err
	}
	return ret, nil
}

// DelKey 删除键值
func DelKey(key string) ( string, error) {
	var Conn = GetRedisPool().Get()
	//defer Conn.Close()
	ret, err := redis.String(Conn.Do("del", key))
	if err!=nil {
		return ret,err
	}
	return ret, nil
}

