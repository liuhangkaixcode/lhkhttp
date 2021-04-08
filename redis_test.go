package lhkhttp

import (
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	aa:=NewRedis()
	aa.SetV("kk",time.Now().String())
}
