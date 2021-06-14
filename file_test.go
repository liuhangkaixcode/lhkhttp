package lhktools

import (
	"testing"
	"time"
)

func TestGetFileContent(t *testing.T) {
	WriteStrTofile("liuhangkai"+time.Now().String()+"\n","111.txt")
}
