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

    SelectMap(sqlStr string)([]map[string]interface{},error)
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
func (s *SqlSt)SelectMap(sqlStr string)([]map[string]interface{},error)  {
	rows, err := s.database.Queryx(sqlStr)
	if err!=nil {
		return nil, err
	}
	var result []map[string]interface{}
	for rows.Next(){
		m := map[string]interface{}{}
		rows.MapScan(m)
		result=append(result,m)
	}
	return result,nil
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



