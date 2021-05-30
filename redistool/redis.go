package redistool

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

	//string操作
	SetV(k,v string)(er error)
	GetV(k string)(string,error)
	//expire 过期时间 0表示不过期  db选择数据0-15 默认值是0
	SetEV(k,v string,expire,db int)(er error)
	GetEV(k string,db int)(string,error)
	//通用操作方法
	CommonHandle(f func(conn redis.Conn))

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
	pass  string
	urlstr string
}

type RedisOption func(s *RedisManger)

func WithPassAndURL(urlstr,pass string) RedisOption {
	return func(s *RedisManger) {
		s.pass=pass
		s.urlstr =urlstr
	}
}

func NewRedis(ops ...RedisOption) RedisIF {
	redismanger =new(RedisManger)
	for _,op:=range ops{
		op(redismanger)
	}
	if len(redismanger.urlstr) == 0{
		redismanger.urlstr ="127.0.0.1:6379"
	}
	var dialOPS  []redis.DialOption

	if len(redismanger.pass) !=0 {
		opspass:=redis.DialPassword(redismanger.pass)
		dialOPS=append(dialOPS,opspass)
	}
	opstimeout:=redis.DialConnectTimeout(time.Second*30)
	dialOPS=append(dialOPS,opstimeout)

	redisonce.Do(func() {
		pool := &redis.Pool{
			MaxIdle:     10,
			MaxActive:   20000,
			IdleTimeout: 10 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", redismanger.urlstr,dialOPS...)
			},
		}

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
func (r *RedisManger) SetV(k,v string)(er error) {
	conn := r.pool.Get()
	defer conn.Close()
	_, er= conn.Do("SET", k, v)
	return
}

func (r *RedisManger)GetV(k string)(string,error){
	conn := r.pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("get", k))
	return s,err
}
func (r *RedisManger)SetEV(k,v string,expire,db int)(er error){
	conn := r.pool.Get()
	defer conn.Close()
	er = selectdb(conn, db)
	if er!=nil{
		return
	}
	if expire ==0 {
		_, er= conn.Do("SET", k, v)
		return
	}
	_,er=conn.Do("SETEX", k, expire,v)
	return
}
func (r *RedisManger)GetEV(k string,db int)(string, error){
	conn := r.pool.Get()
	defer conn.Close()
	er := selectdb(conn, db)
	if er!=nil{
		return "",er
	}
	s, er := redis.String(conn.Do("get", k))
	return s,er

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
						fmt.Println("内部的值",string(zhi))
						res<-string(zhi)

					}
				}
			}
		}
	}()
	//退出的标志
	<-exit
}

func selectdb(conn redis.Conn,db int) error{
	if db>0 {
		_, err := conn.Do("SELECT", db)
		return err
	}
	return nil

}

func (r *RedisManger)CommonHandle(f func(conn redis.Conn))  {
	conn := r.pool.Get()
	defer func() {
		conn.Close()
		fmt.Println("-------------isending")
	}()
	f(conn)
}


