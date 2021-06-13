package lhktools

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	fmt.Println(GetUUID())
	fmt.Println(GetSnowflakeId())
	urlencodeStr:=URLencode("dsafsadfunck尼米的没")
	fmt.Println(urlencodeStr)
	fmt.Println(URLDecode(urlencodeStr))

    base64encry:=Base64Encry("刘航凯信息")
	decry, err := Base64Decry(base64encry)
	fmt.Println(base64encry,decry,err)

}

func TestArivie(t *testing.T)  {
	fmt.Println(GetTimeStamp())
	fmt.Println(GetTimeStamp(0))
	fmt.Println(GetTimeStamp(1))
	fmt.Println(GetTimeStampWithStr("2021-06-13 18:27:31"))
	fmt.Println(GetTimeStrWithTimeStamp(1623580052))
    fmt.Println(GetYearMonthToDay(2021,2))

	_, time1 := GetTimeStampWithStr("1987-01-23 00:00:00")
	_, time2 := GetTimeStampWithStr("1987-01-24 00:00:00")
	fmt.Println(time2.Sub(time1))


}
