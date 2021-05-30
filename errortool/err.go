package errortool

import (
	"encoding/json"
	"fmt"
)

const (
	Code_Unknow="CODE4000"

)
const (
	Msg_JsonError="错误信息无法json序列化"
	Msg_ParamError="code或者msg为空"
)

type ErrMgr struct {
	Code string
	Msg string
}

func (e *ErrMgr)Error() string {
	if len(e.Code) ==0 {
		e.Code=Code_Unknow
	}
	if len(e.Msg) ==0 {
		e.Msg=Msg_ParamError
	}
	marshal, _ := json.Marshal(map[string]string{"code":e.Code,"msg":e.Msg})
	return string(marshal)
}

func NewError(code,msg string) error {
	if len(code) == 0 || len(msg) ==0 {
		return fmt.Errorf("code或者msg为空")
	}
	return &ErrMgr{Code: code,Msg: msg}
}

func GetErrInfo(err error) (code,msg string){

	return errInfo(err)

}
func GetErrMap(err error) (map[string]string){
	code,msg:=errInfo(err)
	return map[string]string{"code":code,"msg":msg}
}

func errInfo(err error) (code,msg string)  {
	if v,ok:=err.(*ErrMgr);ok{
		return v.Code,v.Msg
	}else{
		str:=err.Error()
		if len(str) ==0 {
			 return Code_Unknow,"msg信息为空"
		}

		var m map[string]string
		err := json.Unmarshal([]byte(str), &m)
		if err!=nil{
			return Code_Unknow,str
		}
		if len(m["code"]) ==0 || len(m["msg"])==0{
			return Code_Unknow,str
		}
		return m["code"],m["msg"]
	}

	return Code_Unknow,Msg_ParamError
}

func HasErr(e error) bool {
	if e!=nil{
		return true
	}
	return false
}
