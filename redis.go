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
	//string操作 expireTime 过期时间 0表示不过期
	SetV(k,v string,expireTime int)(er error)
	GetV(k string)(string,error)
	//list 入栈操作操作 command (RPUSH,LPUSH)  v 第一个是key 后面依次是value
	LorRPUSH(command string,v...interface{})error
	//POP操作 command(RPOP LPOP)
	LorRPOP(command,k string)(string ,error)
	//设置过期时间
	
	//获取过期时间
	//阻塞式的获取队列 BLPOP BRPOP
	B_L_R_POP(command,k string,idleTime int,stop <-chan int,res chan <-string)




}
type RedisManger struct {
	pool *redis.Pool
}

func NewRedis(url string) RedisIF {
	if len(url) == 0{
		url ="127.0.0.1:6379"
	}
	redisonce.Do(func() {
		pool := &redis.Pool{
			MaxIdle:     10,
			MaxActive:   20000,
			IdleTimeout: 10 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", url,redis.DialConnectTimeout(time.Second*30))
			},
		}
		redismanger=new(RedisManger)
		redismanger.pool=pool
		conn := pool.Get()
		defer conn.Close()

		_, err := conn.Do("ping")
		if err != nil {
			panic("redis server 未启动...\n")
		}else{
			fmt.Println("redis SUCCESS....")
		}

	})
	return redismanger
}
//string操作 expireTime 过期时间 0表示不过期
func (r *RedisManger) SetV(k,v string,expireTime int)(er error) {
	conn := r.pool.Get()
	defer conn.Close()
	if expireTime == 0 {
		_, er= conn.Do("SET", k, v)
	}else{
		_,er=conn.Do("SETEX", k, expireTime,v)
	}
	return
}
func (r *RedisManger)GetV(k string)(string,error){
	conn := r.pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("get", k))
	return s,err
}
func (r *RedisManger)LorRPUSH(command string,v...interface{})error{
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do(command, v...)
	return err
}
func (r *RedisManger)LorRPOP(command,k string)(string ,error){
	conn := r.pool.Get()
	defer conn.Close()
	res1, err := redis.String(conn.Do(command, k))
	if err != nil {
		return "",err
	}
	return res1,nil

}
func (r *RedisManger)B_L_R_POP(command,k string,idleTime int,exit <-chan int,res chan <-string)  {
	conn := r.pool.Get()
	defer func() {
		conn.Close()
		fmt.Println("B_L_R_POP方法已经退出了")
	}()
	if idleTime == 0{
		idleTime =10
	}
	go func() {
		for{
			s, e := redis.Values(conn.Do(command, k, idleTime))
			if e!=nil {
				continue
			}else{
				for index,v:=range s{
					if index ==0{
						continue
					}
					if zhi,ok:=v.([]byte);ok {
						res<-string(zhi)
					}
				}
			}
		}
	}()
	//退出的标志
	<-exit
}

