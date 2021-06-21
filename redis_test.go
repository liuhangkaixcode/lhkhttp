package lhktools

import (
	"fmt"
	"testing"
	"time"
)

//go test -v --race  -run="TestNewRedis"
func TestNewRedis(t *testing.T) {
	redis:= NewRedis("14.116.147.19:8787","liuhangkai*#920")
	fmt.Println(redis.SetV("yy-yy","liuhangkai===>jier"))
}
func TestBPOP(t *testing.T)  {
	defer fmt.Println("=====testBPOP已经退出了")
	redis := NewRedis("14.116.147.19:8787","liuhangkai*#920")
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



