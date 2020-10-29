package lhkhttp

import (
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

//get请求
func TestGet(t *testing.T) {

	c:=NewClient()
	c.Url="http://ip:8787?type=get&name=cccccccccccccccc&score=刘寒假"
	s, e :=c.Get()
	fmt.Print(s,e)
	//{"type":"get","name":"cccccccccccccccc","score":"\u5218\u5bd2\u5047"}
}
//POST form-data请求
func TestPost(t *testing.T) {
	c:=NewClient(WithTimeOut(16))
	c.Url="http://ip:8787?type=post"
	//请求数据体
	c.Data["name"]="刘xxx"
	c.Data["sex"]="cc"
	c.Data["sub"]="xxxxxxxxxxxxxxxx"
	s,e:=c.Post()
	fmt.Print(s,e)
	//{"name":"\u5218xxx","sex":"cc","sub":"xxxxxxxxxxxxxxxx"}
}

//post requestbody
func TestPostBody(t *testing.T) {
	c:=NewClient(WithTimeOut(26))
	c.Url="http://ip:8787?type=postbody"
	c.Data["name"]="刘xxx"
	c.Data["sex"]="cc"
	c.Data["sub"]=map[string]interface{}{"age":10,"sex":"nan","height":172}
	s,e:=c.PostForBody()
	fmt.Print(s,e)

	//{"name":"刘xxx","sex":"cc","sub":{"age":10,"height":172,"sex":"nan"}}

}
//RPCX框架 service http请求
func TestRpcXGateway(t *testing.T)  {
	c:=NewClient(WithTimeOut(5))
	c.Url="http://ip:8972"
    //请求体
	c.Headers["Content-Type"]="application/rpcx"
	c.Headers["X-RPCX-SerializeType"]=1
	c.Headers["X-RPCX-ServicePath"]="service_name_01"
	c.Headers["X-RPCX-ServiceMethod"]="Add"
	//请求数据体
	c.Data["Name"]="xxxxxxxxxxxxx--"
	s,e:=c.PostForBody()
	fmt.Print(s,"错误:",e)

	//{"Stuatus":3033,"Data":"xxxxxxxxxxxxx--server01-server01"}

}


