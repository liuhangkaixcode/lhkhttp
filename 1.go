package lhkhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var(
	method_GET="GET"
	method_POST="POST"

)

func Get(url string)( string, error)  {
    return request(method_GET,url,nil,nil,false)
}
func Post(url string,header map[string]string,data map[string]interface{})( string, error)   {
	return request(method_POST,url,header,data,false)
}

func PostBody(url string,header map[string]string,data map[string]interface{})(string,error)  {
	return request(method_POST,url,header,data,true)
}

func request(method,url string,headers map[string]string,data map[string]interface{},isrequstbody bool) ( string, error) {
	 var body *bytes.Reader
	if isrequstbody {
		if len(data)>0 {
			marshal, _ := json.Marshal(data)
			body=bytes.NewReader(marshal)
		}else{
			body=new(bytes.Reader)
		}
	}else{
		if len(data)>0 {
			body=bytes.NewReader([]byte(getQueryStr(data)))
		}else{
			body=new(bytes.Reader)
		}
	}

		req, e := http.NewRequest(method, url, body)
		if e!=nil {
			return "",e
		}

		for k,v:=range headers{
			req.Header.Set(k,v)
		}

	if len(data)>0 {
		if isrequstbody {
			req.Header.Set("Content-Type","application/json")
		}else{
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}

		client:=&http.Client{}
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