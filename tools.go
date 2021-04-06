package lhkhttp

import (
	"crypto/aes"
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

//aes加密
//go get -u github.com/forgoer/openssl
//http://tool.chacuo.net/cryptrsapubkey


// AesECBEncrypt
func AesECBEncrypt(src, key []byte, padding string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return ECBEncrypt(block, src, padding)
}

// AesECBDecrypt
func AesECBDecrypt(src, key []byte, padding string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return ECBDecrypt(block, src, padding)
}
