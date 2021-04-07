package lhkhttp

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueStr(t *testing.T)  {
	t1:=make(map[string]interface{})
	t1["age"]=10
	t1["year"]="2014"
	t1["age"]=true
	t1["time"]="2012-12"
	fmt.Println(getQueryStr(t1))
	fmt.Println(getQueryStr(map[string]interface{}{}))
	fmt.Println(getQueryStr(map[string]interface{}{"x":"y"}))
}

//普通get请求
func TestGet(t *testing.T) {
	/*第一种方式
	urlstring:="http://ip:8787/get1?type=get&name=玩笑&score=刘寒假"
	fmt.Println(urlstring)
	c:=NewClient()
	s, e :=c.Get(urlstring)
	fmt.Print(s,e)
	*/
    //第二种方式
	/*
	c:=NewClient(WithHost("http://ip:8787"),WithTimeOut(15))
	data:=make(map[string]interface{})
	data["type"]="get"
	data["name"]="张三"
	data["score"]=140.2
	urlstring, err := MapChangeToQueryUrl("/get1", data)
	fmt.Println(urlstring,err)
	result, err := c.Get(urlstring)
	fmt.Println(result,err)
  */
}
//普通的POST 表单提交请求（application/x-www-form-urlencoded）
func TestPost(t *testing.T) {
	//第一种
	//c:=NewClient(WithTimeOut(16),WithHost("http://ip:8787"))
	//datas:=map[string]interface{}{
	//	"name":"李四",
	//	"wangwu":"zhangsan",
	//	"score":100,
	//	"height":109.2,
	//}
	//headers:=map[string]interface{}{
	//	"Secerct":"xxxxwang==",
	//}
	//res, err := c.Post("/post1", datas, headers)
	//fmt.Println(res,err)

    //第二种
	c:=NewClient(WithTimeOut(16),WithHost("http://ip:8787"))
	res, err := c.Post("/post3", nil, nil)
	fmt.Println(res,err)
}
//
//post请求是 body提 （application/json）
func TestPostBody(t *testing.T) {
	c:=NewClient(WithTimeOut(17),WithHost("http://ip:8787"))
	datas:=map[string]interface{}{
		"name":"李四",
		"wangwu":"zhangsan",
		"score":100,
		"height":109.2,
	}
	headers:=map[string]interface{}{
		"Secerctbody":"xxxxwang==",
	}
	res, err := c.PostForBody("/post2", datas, headers)
	fmt.Println(res,err)

	var r map[string]interface{}
	json.Unmarshal([]byte(res),&r)
	fmt.Println(r["height"],r["name"])

}
//RPCX框架 service http请求
func TestRpcXGateway(t *testing.T)  {
	//废弃方法
	//c:=NewClient(WithTimeOut(5))
	//c.Url="http://ip:8972"
   ////请求体
	//c.Headers["Content-Type"]="application/rpcx"
	//c.Headers["X-RPCX-SerializeType"]=1
	//c.Headers["X-RPCX-ServicePath"]="service_name_01"
	//c.Headers["X-RPCX-ServiceMethod"]="Add"
	////请求数据体
	//c.Data["Name"]="xxxxxxxxxxxxx--"
	//s,e:=c.PostForBody()
	//fmt.Print(s,"错误:",e)

}

func TestChangeToQueryUrl(t *testing.T) {
	url:="http://www.baidu.com"
	data:=make(map[string]interface{})
	data["name"]="liuhangkai"
	data["age"]=100
	data["height"]=109.12
	data["niu"]="张三"
	fmt.Println(MapChangeToQueryUrl(url,data))
}

//测试 aes ecb
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

