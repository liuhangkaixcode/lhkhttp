package lhktools

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"time"
)

func GetUUID() string{
	return uuid.New().String()
}


var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
)

func GetSnowflakeId() int64 {
	lastTimeStamp = time.Now().UnixNano() / 1000000
	curTimeStamp := time.Now().UnixNano() / 1000000

	// 同一毫秒
	if curTimeStamp == lastTimeStamp {
		sn++
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sn > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1000000
			lastTimeStamp = curTimeStamp
			sn = 0
		}

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}


	if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}


	if curTimeStamp < lastTimeStamp {
		return 0
	}

	return 0
}

/*
作用是在URL传参时如果直接传中文可能会出问题（对中文参数支持部完善），所以先用 Server.UrlEncode("中文参数");编码
到另外一个页面接受的时候在用Server.UrlDecode("编码后参数一般为 %+ 16进制数的形式");解码获取中文参数
*/
func URLencode(str string) string {
	return url.QueryEscape(str)
	//url.QueryUnescape("")
	//v := url.Values{}
	//v.Encode()

}
func URLDecode(str string)(string ,error) {
	return url.QueryUnescape(str)
}

func Base64Encry(str string)string  {
	 data:= base64.StdEncoding.EncodeToString([]byte(str))
	 return string(data)
}
func Base64Decry(str string)(string,error)  {
	data,er1:= base64.StdEncoding.DecodeString(str)
	return string(data),er1
}

//默认是秒  1 是毫秒
func GetTimeStamp(flag ...int) int64{
	if len(flag)>0{
		k:=flag[0]
		if k ==1 {
			return time.Now().UnixNano()/1000000
		}
	}
	return  time.Now().Unix()
}

func GetTimeStampWithStr(str string) (int64 ,time.Time) {
	//"2018-07-11 15:07:51"
	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
	tt, _ := time.ParseInLocation("2006-1-02 15:04:05", str, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	return tt.Unix(),tt

}

func GetTimeStrWithTimeStamp(stamp int64,formateStr ...string) (string,time.Time){
	fs:="2006-01-02 15:04:05"
	if len(formateStr)>0 {
		fs=formateStr[0]
	}
	fmt.Print(fs)
	tm := time.Unix(stamp, 0)

    return tm.Format(fs),tm
}

// getYearMonthToDay 查询指定年份指定月份有多少天
// @params year int 指定年份
// @params month int 指定月份
func GetYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}




