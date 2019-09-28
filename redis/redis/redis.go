package redis

import (
	"github.com/gomodule/redigo/redis"
)

//https://github.com/dockerfile/redis

var c redis.Conn

type MyRedis struct {
	pool *redis.Pool
}

type RedisConfig struct {
	MaxIdle   int
	MaxActive int
	Addr      string
}

//NewRedis
func NewRedis(config *RedisConfig) *MyRedis {
	return &MyRedis{
		pool: &redis.Pool{
			MaxIdle:   config.MaxIdle,
			MaxActive: config.MaxActive,
			Dial: func() (conn redis.Conn, e error) {
				return redis.Dial("tcp", config.Addr)
			},
		},
	}
}

func (r *MyRedis) RedisPool() (string, error) {
	var (
		err error
	)
	c := r.pool.Get()
	defer c.Close()
	if _, err = c.Do("SET", "mykey", "superWang"); err != nil {
		return "", err
	}
	return redis.String(c.Do("GET", "mykey"))
}
