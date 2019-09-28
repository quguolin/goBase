package redis

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewRedis_100(t *testing.T) {
	redis := NewRedis(&RedisConfig{MaxActive: 100, MaxIdle: 100, Addr: "127.0.0.1:6381"})
	wg := &sync.WaitGroup{}
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func() {
			wg.Done()
			if _, err := redis.RedisPool(); err != nil {
				fmt.Printf("error (%+v)\n\r", err)
			}
		}()
	}
	wg.Wait()
}

func TestNewRedis_500(t *testing.T) {
	redis := NewRedis(&RedisConfig{MaxActive: 500, MaxIdle: 500, Addr: "127.0.0.1:6381"})
	wg := &sync.WaitGroup{}
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func() {
			wg.Done()
			if _, err := redis.RedisPool(); err != nil {
				fmt.Printf("error (%+v)\n\r", err)
			}
		}()
	}
	wg.Wait()
}
