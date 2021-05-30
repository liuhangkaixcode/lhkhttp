package encrytool

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)



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

