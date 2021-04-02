package lhkhttp

import (
	"fmt"
	"net/url"
)

func MapChangeToQueryUrl(urlstring string,data map[string]interface{}) (string,error){

	if data == nil || len(data) ==0  || len(urlstring) ==0{
		return urlstring,fmt.Errorf("url或者data为空")
	}
	urlstring=urlstring+"?"
	for k,v:=range data{
		urlstring+=fmt.Sprintf("%v=%v&",k,v)
	}

	urlstring=urlstring[0:len(urlstring)-1]
	return urlstring,nil

}

func URLencode(str string) string {
	return url.QueryEscape(str)
	//url.QueryUnescape("")
	//v := url.Values{}
	//v.Encode()

}
func URLDecode(str string)(string ,error) {
	return url.QueryUnescape(str)
}