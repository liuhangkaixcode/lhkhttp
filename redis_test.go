package lhktools

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
     //testBPOP()
	redis1 := NewRedis(WithPassAndURL("14.116.147.19:8787","aa123456"))
	//fmt.Println(redis.SetV("yy-yy","liuhangkai"))
	//fmt.Println(redis.SetEV("kk-kk-1","vvv-vv",89,12))
	redis1.CommonHandle(func(conn redis.Conn) {

	})

}

func testBPOP()  {
	defer fmt.Println("=====testBPOP已经退出了")
	redis := NewRedis(WithPassAndURL("127.0.0.1:8787","aa123456"))
	exit:=make(chan int)
	res:=make(chan string)
	var stopflag = 0
	go func() {
		for{
			stopflag ++
			select {
			case tt:=<-res:
				fmt.Println("外部的值",tt)
				case <-time.After(time.Second*5):
					fmt.Println("==已经超时=",stopflag)
			}

			if stopflag ==9 {
				 exit<-1
			}
		}
	}()
	redis.B_L_R_POP("BLPOP","p1",10,exit,res)
}



