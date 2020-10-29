package lhkhttp

import (
	"fmt"
	"testing"
)


func TestGet(t *testing.T) {
	s, e := Get("http://14.116.147.19:8787?type=get&name=cccccccccccccccc&score=刘寒假")
	fmt.Print(s,e)
	//{"resultget":{"type":"get","name":"cccccccccccccccc","score":"\u5218\u5bd2\u5047"}
}

func TestPost(t *testing.T) {
	url:="http://14.116.147.19:8787?type=post"
	datas:=make(map[string]interface{})
	datas["name"]="刘xxx"
	datas["sex"]="cc"
	datas["sub"]="xxxxxxxxxxxxxxxx"
	s,e:=Post(url,nil,datas)
	fmt.Print(s,e)
	//{"resultpost":{"name":"\u5218xxx","sex":"cc","sub":"xxxxxxxxxxxxxxxx"}
}


func TestPostBody(t *testing.T) {
	url:="http://14.116.147.19:8787?type=postbody"
	datas:=make(map[string]interface{})
	datas["name"]="刘xxx"
	datas["sex"]="cc"
	datas["sub"]=map[string]interface{}{"age":10,"sex":"nan","height":172}
	s,e:=PostBody(url,nil,datas)
	fmt.Print(s,e)

	//{"name":"刘xxx","sex":"cc","sub":{"age":10,"height":172,"sex":"nan"}}

}

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
