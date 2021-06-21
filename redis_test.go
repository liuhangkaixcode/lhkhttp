package lhktools

import (
	"fmt"
	"testing"
	"time"
)

//go test -v --race  -run="TestNewRedis"
func TestNewRedis(t *testing.T) {
	redis := NewRedis("14.116.147.19:8787", "liuhangkai*#920")
	fmt.Println(redis.SetV("yy-yy", "liuhangkai==||||=>jier"))
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
