package lhkhttp

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

var (
	redismanger  *RedisManger
	redisonce sync.Once
)
type RedisIF interface {
	SetV(k,v string)error

}
type RedisManger struct {
	pool *redis.Pool
}

func NewRedis() RedisIF {
	redisonce.Do(func() {
		pool := &redis.Pool{
			MaxIdle:     10,
			MaxActive:   20000,
			IdleTimeout: 10 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "127.0.0.1:6379",redis.DialConnectTimeout(time.Second*10))
			},
		}
		redismanger=new(RedisManger)
		redismanger.pool=pool
		conn := pool.Get()
		defer conn.Close()

		_, err := conn.Do("ping")
		if err != nil {
			panic("redis is not start........\n")
		}else{
			fmt.Println("redis inint success....")
		}

	})
	return redismanger

}

func (r *RedisManger) SetV(k,v string)error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", k, v)
	return err
}

