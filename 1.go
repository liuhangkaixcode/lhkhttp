package lhkhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var(
	method_GET="GET"
	method_POST="POST"
	method_PUT="PUT"
	method_DELETE="DELETE"
)
type ClientReq struct {
	 timeOut  time.Duration //默认50秒的超时时间
	 method string
	 Data map[string]interface{}     //请求数据体
	 Headers map[string]interface{}  //header数据
	 Url string         //请求url
	 IsRequstbody bool  //是否是requestBody请求
}

type OpFunc  func(c *ClientReq)

func WithTimeOut(t int64) OpFunc  {
	return func(c *ClientReq) {
		c.timeOut=time.Duration(t)*time.Second
	}
}

//初始化
func NewClient(ops...OpFunc) *ClientReq{
	client:=  &ClientReq{Data:make(map[string]interface{}),Headers:make(map[string]interface{})}
	client.timeOut=50*time.Second
	for _,op:=range ops{
		op(client)
	}
	return client
}

func (c *ClientReq)Get()( string, error)  {
	c.method=method_GET
	c.Data=nil
	c.Headers=nil
	c.IsRequstbody=false
    return c.request()
}

func (c *ClientReq)Post()( string, error)   {
	c.method=method_POST
	c.IsRequstbody=false
	return c.request()
}

func (c *ClientReq)PostForBody()(string,error)  {
	c.method=method_POST
	c.IsRequstbody=true
	return c.request()
}

func (c *ClientReq)Put()( string, error)   {
	c.method=method_POST
	c.IsRequstbody=false
	return c.request()
}

func (c *ClientReq)PutForBody()(string,error)  {
	c.method=method_PUT
	c.IsRequstbody=true
	return c.request()
}

func (c *ClientReq)Del()( string, error)   {
	c.method=method_DELETE
	c.IsRequstbody=false
	return c.request()
}

func (c *ClientReq)DelForBody()(string,error)  {
	c.method=method_DELETE
	c.IsRequstbody=true
	return c.request()
}

func (c *ClientReq)request()( string, error) {

	if len(c.Url) == 0 {
		return "",fmt.Errorf("url没有传")
	}

	 var body *bytes.Reader
	if c.IsRequstbody {
		if len(c.Data)>0 {
			marshal, _ := json.Marshal(c.Data)
			body=bytes.NewReader(marshal)
		}else{
			body=new(bytes.Reader)
		}
	}else{
		if len(c.Data)>0 {
			body=bytes.NewReader([]byte(getQueryStr(c.Data)))
		}else{
			body=new(bytes.Reader)
		}
	}

		req, e := http.NewRequest(c.method, c.Url, body)
		if e!=nil {
			return "",e
		}



	if len(c.Data)>0 {
		if c.IsRequstbody {
			req.Header.Set("Content-Type","application/json")
		}else{
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	for k,v:=range c.Headers{
		req.Header.Set(k,fmt.Sprintf("%v",v))
	}

		client:=&http.Client{}
		fmt.Println("超时时间",c.timeOut)
		client.Timeout=c.timeOut
		response, e := client.Do(req)
		if e!=nil {
			return  "",e
		}
		defer response.Body.Close()

		bytes, e := ioutil.ReadAll(response.Body)
		if e != nil {
			return "", e
		}
		return string(bytes),nil


}

//post "k=v&k2=v2"  postbody {"k":"v","k2":"v2"}
func getQueryStr(datas map[string]interface{}) string  {
	if len(datas) ==0 {
		return  ""
	}

	result:=""

	for k,v:=range datas{
		result=fmt.Sprintf("%s&%s=%v",result,k,v)
	}

	return result[1:]

}