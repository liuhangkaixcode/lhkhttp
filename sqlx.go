package lhkhttp

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"

)
var (
	once sync.Once
)

type SqlIF interface {

}
func NewMysql()  SqlIF{
	once.Do(func() {

		fmt.Print("=========once")
	})
	fmt.Println("=======")

}
func InitMySql() (*sqlx.DB,error) {
	dns := "root:liuhangkai*#920@tcp(159.75.42.138:3306)/test1"
	database, err := sqlx.Connect("mysql", dns)
	return database,err

}
