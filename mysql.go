package lhktools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	sqlmanger *SqlManger
)

type SqlIF interface {
	//获取数据实例
	//GetDBInstance() *sqlx.DB
	//关闭
	Close()
	//单行数据 sqlstr sql查询语句  obj 结构体对象  args查询语句可选参数
	Get(sqlStr string, obj interface{}, args ...interface{}) error
	//多行
	Select(sqlStr string, objs interface{}, args ...interface{}) error
	//返回多行map
	SelectMap(sqlStr string) ([]map[string]interface{}, error)
	//insert
	Insert(sqlStr string, args ...interface{}) (insertId int64, err error)
	//updateOrDelete
	UpdateOrDelete(sqlStr string, args ...interface{}) (rowsAffect int64, err error)
	//事务操作
	BeginHandle(f func(tx *sql.Tx, err error) error)
	//事务查询
	BeginQuery(tx *sql.Tx, sqlStr string, args ...interface{}) error
	//事务写操作
	BeginExec(tx *sql.Tx, sqlStr string, args ...interface{}) error
}
type SqlManger struct {
	database *sqlx.DB
}

func (s *SqlManger) Get(sqlstr string, obj interface{}, args ...interface{}) error {
	return s.database.Get(obj, sqlstr, args...)
}
func (s *SqlManger) Select(sqlstr string, objs interface{}, args ...interface{}) error {
	return s.database.Select(objs, sqlstr, args...)
}
func (s *SqlManger) SelectMap(sqlStr string) ([]map[string]interface{}, error) {
	rows, err := s.database.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	for rows.Next() {
		m := map[string]interface{}{}
		rows.MapScan(m)
		result = append(result, m)
	}
	return result, nil
}
func NewMysql(dns string) SqlIF {

	if len(dns) == 0 {
		panic("mysql连接服务器的地址没有传")
	}
	//dns := "root:liuhangkai*#920@tcp(159.75.42.138:3306)/test1"
	database, err := sqlx.Connect("mysql", dns)
	if err != nil {
		panic("mysql服务器无法连接")
	} else {
		fmt.Println("mysql正常启动")
	}
	sqlmanger = new(SqlManger)
	sqlmanger.database = database

	return sqlmanger
}

func (s *SqlManger) Insert(sqlStr string, args ...interface{}) (insertId int64, err error) {
	exec, err := s.database.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}
	insertId, err = exec.LastInsertId()
	return
}

func (s *SqlManger) UpdateOrDelete(sqlStr string, args ...interface{}) (rowsAffect int64, err error) {
	exec, err := s.database.Exec(sqlStr, args...)
	if err != nil {
		return 0, err
	}
	rowsAffect, err = exec.RowsAffected()
	return
}

//func (s *SqlManger)GetDBInstance() *sqlx.DB{
//	return s.database
//}
func (s *SqlManger) Close() {
	if sqlmanger != nil {
		sqlmanger.database.Close()
	}
}

//事务操作
func (s *SqlManger) BeginHandle(f func(t *sql.Tx, err error) error) {
	tx, e := s.database.Begin()
	err := f(tx, e)
	if e != nil {
		return
	}
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}

func (s *SqlManger) BeginQuery(tx *sql.Tx, sqlstr string, args ...interface{}) error {
	rr := tx.QueryRow(sqlstr)
	err := rr.Scan(args...)
	return err
}

func (s *SqlManger) BeginExec(tx *sql.Tx, sqlstr string, args ...interface{}) error {
	_, err := tx.Exec(sqlstr, args...)
	return err
}
