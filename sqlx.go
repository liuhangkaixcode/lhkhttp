package lhkhttp

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)
var (
	once sync.Once
	sql *SqlSt
)

type SqlIF interface {
    Close()
    Query(str string)
}
type SqlSt struct {
	database *sqlx.DB
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
		}

		sql= new(SqlSt)
		sql.database=database

	})

   return sql
}

func (s *SqlSt)Close () {
	if sql !=nil{
		sql.Close()
	}
}



