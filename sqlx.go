package lhkhttp

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)
var (
	once sync.Once
	sqlhandle *SqlManger
)

type SqlIF interface {
    Close()
    //单行数据
    Get(sqlstr string,obj interface{})(error)
    //多行
    Select(sqlStr string,objs interface{})(error)
    //多行map
    SelectMap(sqlStr string)([]map[string]interface{},error)
    //insert
    Insert(sqlStr string)(insertId int64,err error)
    //updateOrDelete
	UpdateOrDelete(sqlStr string)(rowsAffect int64,err error)
}
type SqlManger struct {
	database *sqlx.DB
}

func (s *SqlManger)Get(sqlstr string,obj interface{}) (error) {
   return s.database.Get(obj,sqlstr)
}
func (s *SqlManger)Select(sqlstr string,objs interface{}) (error) {
	exec, _ := s.database.Exec("")
	exec.LastInsertId()
	exec.RowsAffected()
	return  s.database.Select(objs,sqlstr)
}
func (s *SqlManger)SelectMap(sqlStr string)([]map[string]interface{},error)  {
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

		sqlhandle= new(SqlManger)
		sqlhandle.database=database

	})

   return sqlhandle
}

func (s *SqlManger)Insert(sqlStr string)(insertId int64,err error){
	exec, err := s.database.Exec(sqlStr)
	if err!=nil {
		return 0, err
	}
	insertId,err= exec.LastInsertId()
	return
}

func (s *SqlManger)UpdateOrDelete(sqlStr string)(rowsAffect int64,err error){
	exec, err := s.database.Exec(sqlStr)
	if err!=nil {
		return 0, err
	}
	rowsAffect,err= exec.RowsAffected()
	return
}

func (s *SqlManger)Close () {
	if sqlhandle !=nil{
		sqlhandle.database.Close()


	}
}



