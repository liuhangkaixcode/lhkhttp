package lhkhttp

import (
	"database/sql"
	"fmt"
	"testing"
)
//name "name": converting NULL to string is unsupported  ==>sql.NullString
type Stu struct {
	ID int `db:"id""`
	Name sql.NullString `db:"name"`
	Total int  `db:"total"`
	Birth string `db:"birth"`
}
//数据库
var(
  dns1 = "root:123456@tcp(14.116.147.19:8787)/skill?timeout=10s&readTimeout=12s"
    sqlstrucet = NewMysql(dns1)
)



func TestInitMySql(t *testing.T) {

	//单行
	//signlesqlstr:=fmt.Sprintf("select id,name,total,birth from stu where id='%d' ",1)
	//var stu Stu
	//err := sqlstrucet.Get(signlesqlstr, &stu)
	//if err!=nil {
	//	fmt.Print(err)
	//}else{
	//	marshal, _ := json.Marshal(stu)
	//	fmt.Println("==>",string(marshal))
	//	fmt.Println(stu.Name.String)
	//}

	//获取一个属性
	siglenamesql:=fmt.Sprintf("select name from stu where id='%d'",1)
	var name string
	err := sqlstrucet.Get(siglenamesql, &name)
	fmt.Println(err,name)

	//多行
	//doublesqlstr:=fmt.Sprintf("select id,name,total from stu")
	//var stus []Stu
	//err := sqlstrucet.Select(doublesqlstr, &stus)
	//if err!=nil {
	//	fmt.Print(err)
	//}else{
	//	marshal, _ := json.Marshal(stus)
	//	fmt.Println("==>",string(marshal))
	//}

	//多行map[string]interface{}返回
	//doublesqlstr:=fmt.Sprintf("select id,name,total from stu")
	//result,_ := sqlstrucet.SelectMap(doublesqlstr)
	//for _,m:=range result{
	//	for k,v:=range m{
	//		if k=="id" || k == "total"{
	//			atoi, _ := strconv.Atoi(string(v.([]byte)))
	//			m[k]=atoi
	//		}else{
	//			m[k]=string(v.([]byte))
	//		}
	//
	//	}
	//}
	//marshal, _ := json.Marshal(result)
	//fmt.Println("==>",string(marshal))



}

