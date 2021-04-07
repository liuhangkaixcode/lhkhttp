package lhkhttp

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
)
//name "name": converting NULL to string is unsupported  ==>sql.NullString
type Stu struct {
	ID int `db:"id""`
	Name sql.NullString `db:"name"`
	Total int  `db:"total"`
}
//数据库
var(
  dns1 = "root:123456@tcp(14.116.147.19:8787)/skill?timeout=10s&readTimeout=12s"
    sqlstrucet = NewMysql(dns1)
)



func TestInitMySql(t *testing.T) {

	//单行
	signlesqlstr:=fmt.Sprintf("select * from stu where id='%d' ",1)
	fmt.Println(signlesqlstr)
	var stu Stu
	err := sqlstrucet.Get(signlesqlstr, &stu)
	if err!=nil {
		fmt.Print(err)
	}else{
		marshal, _ := json.Marshal(stu)
		fmt.Println("==>",string(marshal))
		fmt.Println(stu.Name.String)
	}


	//多行
	//doublesqlstr:=fmt.Sprintf("select * from stu")
	//var stus []Stu
	//err := sqlstrucet.Select(doublesqlstr, &stus)
	//if err!=nil {
	//	fmt.Print(err)
	//}else{
	//	marshal, _ := json.Marshal(stus)
	//	fmt.Println("==>",string(marshal))
	//}


}

