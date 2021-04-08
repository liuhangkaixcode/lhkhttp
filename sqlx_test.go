package lhkhttp

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)
//name "name": converting NULL to string is unsupported  ==>sql.NullString
//乐观锁
//Select version  from xx where id=1   ++version=4
//update xx set store=store+1,version=version+1 where version=4 and id=1
//测试事务
/*
	*r := tx.QueryRow("select id,name,total from stu where id=1 for update")
			var id int
			var name string
			var total int
			rr.Scan(&id, &name, &total)
	 tx.Exec("insert into or1 (name) values(?)", time.Now().String()+fmt.Sprintf("==>%v", ipaddr))
*/
type Stu struct {
	ID int `db:"id""`
	Name sql.NullString `db:"name"`
	Total int  `db:"total"`
	Birth string `db:"birth"`
}
//数据库
var(
  dns1 = "root:123456@tcp(ip:8787)/skill?timeout=10s&readTimeout=12s"
    sqlstrucet = NewMysql(dns1)

)
//查询
func TestInitMySql(t *testing.T) {
     //defer  sqlstrucet.Close()  //销毁连接
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
	siglenamesql:=fmt.Sprintf("select birth from stu where id='%d'",1)
	var name string
	err := sqlstrucet.Get(siglenamesql, &name)
	fmt.Println(err,name)

	//时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", name, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix(),tt.Year())                             //1531292871

	//时间戳 to 时间
	tm := time.Unix(tt.Unix(), 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19

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

//插入
func TestSqlManger_Insert(t *testing.T) {
	//插入
	//sqlstr:=fmt.Sprintf("insert into stu(name,total)values('%s','%d')","刘航军",123)
	//fmt.Println(sqlstrucet.Insert(sqlstr))

    //更新
	//sqlstr:=fmt.Sprintf("update stu set name='%s' where id<'%d'","更新刘航军名字1",3)
	//fmt.Println(sqlstrucet.UpdateOrDelete(sqlstr))

	//删除
	//sqlstr:=fmt.Sprintf("delete from stu where id='%d'",1)
	//fmt.Println(sqlstrucet.UpdateOrDelete(sqlstr))


    //事务操作
	sqlstrucet.BeginHandle(func(tx *sql.Tx,er error) error {
		if er!=nil{
			return fmt.Errorf("开启事务失败")
		}
		//开始执行业务逻辑
		var id int
		var name string
		var total int
		er = sqlstrucet.BeginQuery(tx, "select id,name,total form stu where id=1 for update", &id, &name, &total)
		fmt.Println(id,name,total,er)
		if er!=nil {
			return er
		}
       return nil
	})
	fmt.Println("=======ceshi")




}

