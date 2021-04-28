package lhkhttp

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
     //testBPOP()
}

func testBPOP()  {
	defer fmt.Println("=====testBPOP已经退出了")
	redis := NewRedis("")
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



