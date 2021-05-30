package encrytool

import (
	"encoding/base64"
	"fmt"
	"testing"
)

//测试 aes ecb
//http://tool.chacuo.net/cryptdes
func TestAesEncrypt(t *testing.T) {
	//加密
	jsonstr:=`{"w":"w xx w "}`
	src := []byte(jsonstr) //原明文
	// AES-128-ECB, PKCS7_PADDING 输出base64 utf-8编码
	key := []byte("a$efkghm@hkybu#%") //16位密码
	dst, _ := AesECBEncrypt(src, key, PKCS7_PADDING)
	str:=base64.StdEncoding.EncodeToString(dst)
	fmt.Println(str)
	//解密
	src, _ = base64.StdEncoding.DecodeString(str)
	dst, _ = AesECBDecrypt(src, key, PKCS7_PADDING)
	fmt.Println(string(dst))
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5Str("kaishao")) //32位小写
	fmt.Println(Sha1ToStr("kaishao"))

}

