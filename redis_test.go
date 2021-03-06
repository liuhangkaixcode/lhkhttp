package lhktools

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
	"time"
)

//go test -v --race  -run="TestNewRedis"
func TestNewRedis(t *testing.T) {
	//pool 允许重新设置值
	redis := NewRedis("14.116.147.19:8787", "liuhangkai*#920", func(pool interface{}) {
		if v,ok:=pool.(*redis.Pool);ok{
			v.MaxIdle=100
		}
	})
	fmt.Println(redis.SetV("yy-yy", "liuhangkai==||||=>jier"+time.Now().String()))
}
func TestRedisManger_RedisCommonHandle(t *testing.T) {
	red := NewRedis("14.116.147.19:8787", "liuhangkai*#920")
	red.RedisCommonHandle(func(connect interface{}) {
		if conn,ok:=connect.(redis.Conn);ok{
			fmt.Println(redis.Int64(conn.Do("EXISTS","yy4-yy")))
		}
	})
}

func TestRedisManger_HGET(t *testing.T) {
	red := NewRedis("14.116.147.19:8787", "liuhangkai*#920")
	//red.HSET("person","name","liuhangkai")
	//red.HSET("person","age","10")
	//red.HSET("person","height",172.12)
	//fmt.Println(red.HGET("person","height1"))
	fmt.Println(red.HGETALL("person"))
	fmt.Println(red.EXISTS("person1"))
}
func TestBPOP(t *testing.T) {
	defer fmt.Println("=====testBPOP已经退出了")
	redis := NewRedis("14.116.147.19:8787", "liuhangkai*#920")
	exit := make(chan int, 1)
	var stopflag = 0
	res := redis.B_L_R_POP("BLPOP", "p1", 10, exit)
	for {
		stopflag++
		select {
		case tt := <-res:
			fmt.Println("外部的值", tt)
		case <-time.After(time.Second * 3):
			fmt.Println("==已经超时=", stopflag)
		}
		if stopflag == 3 {
			exit <- 1
			time.Sleep(time.Second * 10)
			return
		}
	}
}
