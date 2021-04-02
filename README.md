# lhkhttp  httpClient的封装

```
//发送普通get请求
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
//普通的POST 表单提交请求
func TestPost(t *testing.T) {
	//第一种 带参数
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

    //第二种 不带参数的
	c:=NewClient(WithTimeOut(16),WithHost("http://ip:8787"))
	res, err := c.Post("/post3", nil, nil)
	fmt.Println(res,err)
}
//
//post请求 requestBody方式
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
	
    //数据解析
	var r map[string]interface{}
	json.Unmarshal([]byte(res),&r)
	fmt.Println(r["height"],r["name"])

}
```
