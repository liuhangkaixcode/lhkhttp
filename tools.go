package lhkhttp

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
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

// Md5 Calculate the md5 hash of a string
func Md5(str string) []byte {
	h := md5.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

func Md5Str(src string) string {
	dst := Md5(src)
	return  hex.EncodeToString(dst)
}

func Sha1ToStr(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//RSA  https://www.cnblogs.com/zhichaoma/p/12516715.html
//RSA2  http://www.manongjc.com/detail/16-vxioingzzhrekpa.html

