package mysql

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"sync"
)

var DB *sql.DB
var once sync.Once

func InitMysql(msql string) {
	once.Do(func(){
		var err error
		DB, err = sql.Open("mysql", msql)
		if err != nil {
			beego.Error("连接数据库 【失败】")
			panic(err)
		}
		beego.Info("连接数据库 【完成】")
	})
}

func CreateTable(m *DBModel) error{
	stmt, err := DB.Prepare(m.TableSql())
	if err != nil {
		beego.Error(err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		beego.Error(err)
		return err
	}
	beego.Info("创建表 ",m.Name," 成功 【完成】")
	return nil
}

func Add(m *DBModel) error{

	stmt, err := DB.Prepare(m.InsertSql())
	if err != nil {
		beego.Error("44 ",err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.InsertArgs()...)
	if err != nil {
		beego.Error(err)
		return err
	}
	beego.Info("插入数据成功")
	return nil
}

