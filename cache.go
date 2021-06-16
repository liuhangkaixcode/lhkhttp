package lhktools
import (
"fmt"
"strconv"
"sync"
"time"
)
var (
	myonce sync.Once
	obj *CacheManger
	cacheMap=make(map[string]map[string]string ,1000)
)
type CacheIF interface {
	SetCache(k,v string,inter int)
	GetValue(k string)string
	GetCount() int
	ClearALL()
}
type CacheManger struct {
	clearData int
	sync.RWMutex
}

func (c *CacheManger)SetCache(k,v string,expire int)  {
	if len(k) ==0 || len(v) ==0{
		return
	}
	if expire<2{
		expire = 10
	}
	c.Lock()
	temp:=make(map[string]string)
	sss:=fmt.Sprintf("%d",time.Now().Unix())
	temp["timestamp"]=sss
	temp["expiretime"]=fmt.Sprintf("%d",expire)
	temp["value"]=v
	cacheMap[k]=temp
	c.Unlock()

}
func (c *CacheManger)GetValue(k string)string  {
	if len(k) == 0 {
		return ""
	}
	c.Lock()
	defer c.Unlock()
	if v,ok:=cacheMap[k];ok{
		timenow:=time.Now().Unix()
		timestamp,_:=strconv.Atoi(v["timestamp"])
		expireTime,_:=strconv.Atoi(v["expiretime"])
		value,_:=v["value"]
		delete(cacheMap,k)
		if timenow-int64(timestamp)<int64(expireTime){
			return value
		}
		return ""
	}
	return ""
}
func (c *CacheManger)GetCount() int  {
	c.RLock()
	defer c.RUnlock()
	return len(cacheMap)
}
func (c *CacheManger)ClearALL()  {

}
func NewCache()  CacheIF{
	myonce.Do(func() {
		obj=new(CacheManger)
		fmt.Println("====初始化一次")
		go clearData(obj)
	})
	return obj
}
func clearData(obj *CacheManger) {
	timer:=time.NewTicker(time.Second*5)
	for _= range timer.C{
		obj.Lock()
		for k,v:=range cacheMap{
			temp, _ := strconv.Atoi(v["timestamp"])
			if time.Now().Unix() - int64(temp)>10{
				fmt.Println("key 为",k,"已经删除")
				delete(cacheMap,k)
			}
		}
		obj.Unlock()
	}
}

