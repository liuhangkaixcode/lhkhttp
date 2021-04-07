package lhkhttp

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)
var (
	once sync.Once
	sqlhandle *SqlSt
)

type SqlIF interface {
    Close()
    //单行数据
    Get(sqlstr string,obj interface{})(error)
    //多行
    Select(sqlStr string,objs interface{})(error)
}
type SqlSt struct {
	database *sqlx.DB
}

func (s *SqlSt)Get(sqlstr string,obj interface{}) (error) {
   return s.database.Get(obj,sqlstr)
}
func (s *SqlSt)Select(sqlstr string,objs interface{}) (error) {
    return  s.database.Select(objs,sqlstr)
}
func NewMysql(dns string)  SqlIF{

	once.Do(func() {
		if len(dns)==0 {
			panic("mysql连接服务器的地址没有传")
		}
		//dns := "root:liuhangkai*#920@tcp(159.75.42.138:3306)/test1"
		database, err := sqlx.Connect("mysql", dns)
		if err!=nil {
			panic("mysql服务器无法连接")
		}else{
			fmt.Println("mysql正常启动")
		}

		sqlhandle= new(SqlSt)
		sqlhandle.database=database

	})

   return sqlhandle
}

func (s *SqlSt)Close () {
	if sqlhandle !=nil{
		sqlhandle.database.Close()
	}
}



