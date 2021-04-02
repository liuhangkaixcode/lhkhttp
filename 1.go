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
	 data map[string]interface{}     //请求数据体
	 headers map[string]interface{}  //header数据
	 suburl string         //请求url
	 isRequstbody bool  //是否是requestBody请求
	 host string   //请求服务器的域名地址 http://www.baidu.com
}

type OpFunc  func(c *ClientReq)

func WithTimeOut(t int64) OpFunc  {
	return func(c *ClientReq) {
		c.timeOut=time.Duration(t)*time.Second
	}
}
func WithHost(host string)OpFunc  {
	return func(c *ClientReq) {
		c.host=host
	}
}

//初始化
func NewClient(ops...OpFunc) *ClientReq{
	client:=  &ClientReq{
				data:make(map[string]interface{}),
				headers:make(map[string]interface{})}
				client.timeOut=50*time.Second
				for _,op:=range ops{
				op(client)
		        }
	return client
}

func (c *ClientReq)Get(suburl string)(string, error)  {
	c.method=method_GET
	c.suburl=suburl
	c.data=nil
	c.headers=nil
	c.isRequstbody=false
    return c.request()
}

func (c *ClientReq)Post(suburl string,data,headers map[string]interface{} )( string, error)   {
	c.method=method_POST
	c.isRequstbody=false
	c.suburl=suburl
	c.data=data
	c.headers=headers
	return c.request()
}

func (c *ClientReq)PostForBody(suburl string,data,headers map[string]interface{})(string,error)  {
	c.method=method_POST
	c.isRequstbody=true
	c.suburl=suburl
	c.data=data
	c.headers=headers
	return c.request()
}

func (c *ClientReq)request()( string, error) {
    //构建request请求
	if len(c.suburl) == 0 {
		return "",fmt.Errorf("suburl没有传")
	}
	if len(c.host) !=0 {
		c.suburl=c.host+c.suburl
	}
	 var body *bytes.Reader
	body=new(bytes.Reader)
	if c.isRequstbody {
		if len(c.data)>0 {
			marshal, _ := json.Marshal(c.data)
			body=bytes.NewReader(marshal)
		}
	}else{
		if len(c.data)>0 {
			body=bytes.NewReader([]byte(getQueryStr(c.data)))
		}
	}

	req, e := http.NewRequest(c.method, c.suburl, body)
	if e!=nil {
		return "",e
	}

	if len(c.data)>0 {
		if c.isRequstbody {
			req.Header.Set("Content-Type","application/json")
		}else{
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		for k,v:=range c.headers{
			req.Header.Set(k,fmt.Sprintf("%v",v))
		}
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

