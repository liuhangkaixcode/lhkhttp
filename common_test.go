package lhktools

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	fmt.Println(GetUUID())
	fmt.Println(GetSnowflakeId())
	urlencodeStr:=URLencode("aaafdf=abcdefehttp://cdsdf")
	fmt.Println(urlencodeStr)
	fmt.Println(URLDecode(urlencodeStr))
}
