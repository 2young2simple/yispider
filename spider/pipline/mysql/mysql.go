package mysql

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var DB *sql.DB

func InitMysql(mysql string) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysql)
	orm.RegisterModel(&C{})
	orm.RunSyncdb("default", false, true)
}

type C struct {
	Id int
}

func CreateTable(m *DBModel) error{
	o := orm.NewOrm()
	_,err := o.Raw(m.TableSql()).Exec()
	if err != nil{
		return err
	}
	beego.Info("创建表 ",m.Name," 成功 【完成】")
	return nil
}

func Add(m *DBModel) error{
	o := orm.NewOrm()
	_,err := o.Raw(m.InsertSql(),m.InsertArgs()...).Exec()
	if err != nil{
		return err
	}
	beego.Info("插入数据成功")
	return nil
}

