package lhktools

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redismanger *RedisManger
)

type RedisIF interface {
	//判断某个key是否存在
	EXISTS(k string)(bool)
	//string操作
	SetV(k, v string) (er error)
	GetV(k string) (string, error)
	//expire 过期时间 0表示不过期  db选择数据0-15 默认值是0
	SetEV(k, v string, expire, db int) (er error)
	GetEV(k string, db int) (string, error)
	//允许自己扩充自己的通用操作方法 connect 是redis.Conn类型
	RedisCommonHandle(f func(connect interface{}))
	//list 入栈操作操作 command (RPUSH,LPUSH)  v 第一个是key 后面依次是value
	LorRPUSH(command string, v ...interface{}) error
	//POP操作 command(RPOP LPOP)
	LorRPOP(command, k string) (string, error)
	//阻塞式的获取队列 BLPOP BRPOP
	B_L_R_POP(command, k string, idleTime int, exit chan int) (res chan string)
	//map操作 v的值 {person:{k1,v1}}
	HMSET(v ...interface{}) error
	HSET(v ...interface{}) error
	HGET(key, field string) (string, error)
	HGETALL(key string) (map[string]string, error)
}
type RedisManger struct {
	pool       *redis.Pool
	pwd        string
	connAdress string
}

type RedisOption func(s *RedisManger)

//func WithPassAndURL(urlstr,pass string) RedisOption {
//	return func(s *RedisManger) {
//		s.pass=pass
//		s.urlstr =urlstr
//	}
//
//	for _,op:=range ops{
//		op(redismanger)
//	}
//}

func NewRedis(connAdress, pwd string, f ...func(pool interface{})) RedisIF {
	redismanger = new(RedisManger)
	redismanger.connAdress = connAdress
	redismanger.pwd = pwd
	if len(redismanger.connAdress) == 0 {
		redismanger.connAdress = "127.0.0.1:6379"
	}
	var dialOPS []redis.DialOption

	if len(redismanger.pwd) != 0 {
		opspass := redis.DialPassword(redismanger.pwd)
		dialOPS = append(dialOPS, opspass)
	}
	opstimeout := redis.DialConnectTimeout(time.Second * 30)
	dialOPS = append(dialOPS, opstimeout)

	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   20000,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redismanger.connAdress, dialOPS...)
		},
	}
	if len(f) > 0 {
		f[0](pool)
	}

	redismanger.pool = pool
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("ping")
	if err != nil {
		fmt.Println("===>error", err)
		panic("redis server 未启动...\n")
	} else {
		fmt.Println("redis SUCCESS....")
	}

	return redismanger
}


func (r *RedisManger) EXISTS(k string)(bool) {
	conn := r.pool.Get()
	defer conn.Close()
	i, err := redis.Int64(conn.Do("EXISTS", k))
	if err!=nil{
		return false
	}
	if i == 1{
		return true
	}
	return false
}

//string操作 expireTime 过期时间 0表示不过期
func (r *RedisManger) SetV(k, v string) (er error) {
	conn := r.pool.Get()
	defer conn.Close()
	_, er = conn.Do("SET", k, v)
	return
}

func (r *RedisManger) GetV(k string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("get", k))
	return s, err
}
func (r *RedisManger) SetEV(k, v string, expire, db int) (er error) {
	conn := r.pool.Get()
	defer conn.Close()
	er = selectdb(conn, db)
	if er != nil {
		return
	}
	if expire == 0 {
		_, er = conn.Do("SET", k, v)
		return
	}
	_, er = conn.Do("SETEX", k, expire, v)
	return
}
func (r *RedisManger) GetEV(k string, db int) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	er := selectdb(conn, db)
	if er != nil {
		return "", er
	}
	s, er := redis.String(conn.Do("get", k))
	return s, er

}

func (r *RedisManger) LorRPUSH(command string, v ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do(command, v...)
	return err
}

func (r *RedisManger) HMSET(v ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("HMSET", v...)
	return err
}
func (r *RedisManger) HSET(v ...interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", v...)
	return err
}

func (r *RedisManger) HGET(key, field string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	s, er := redis.String(conn.Do("hget", key, field))
	return s, er
}
func (r *RedisManger) HGETALL(key string) (map[string]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	resultMap:=make(map[string]string)
	for i:=0;i<len(result);i+=2{
		k:=string(result[i].([]byte))
		v:=string(result[i+1].([]byte))
		resultMap[k]=v
	}
	return resultMap, nil

}
func (r *RedisManger) LorRPOP(command, k string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	res1, err := redis.String(conn.Do(command, k))
	if err != nil {
		return "", err
	}
	return res1, nil

}
func (r *RedisManger) B_L_R_POP(command, k string, idleTime int, exit chan int) (res chan string) {
	conn := r.pool.Get()

	res = make(chan string, 100)
	if idleTime == 0 {
		idleTime = 10
	}
	go func(conn redis.Conn, exit chan int) {
		for {

			select {
			case <-exit:
				conn.Close()
				//fmt.Println("===conn已经退出了")
				return
			default:
				{
					s, e := redis.Values(conn.Do(command, k, idleTime))
					if e != nil {
						continue
					} else {
						for index, v := range s {
							if index == 0 {
								continue
							}
							if zhi, ok := v.([]byte); ok {
								//fmt.Println("内部的值", string(zhi))
								res <- string(zhi)

							}
						}
					}
				}

			}

		}
	}(conn, exit)
	//退出的标志
	return
}

func selectdb(conn redis.Conn, db int) error {
	if db > 0 {
		_, err := conn.Do("SELECT", db)
		return err
	}
	return nil

}

func (r *RedisManger) RedisCommonHandle(f func(connect interface{})) {
	conn := r.pool.Get()

	defer func() {
		conn.Close()
		//fmt.Println("-------------isending")
	}()
	f(conn)
}
