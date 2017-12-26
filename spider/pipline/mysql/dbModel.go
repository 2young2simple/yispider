package mysql

import (
	"fmt"
	"strings"
	"reflect"
	"encoding/json"
)

type Field struct {
	Name string
	Pk   bool
	Value interface{}
}

func (f *Field) Sql() string{
	var sql string
	switch f.Value.(type) {
	case string:
		sql = fmt.Sprintf("\n `%s` varchar(255) NULL DEFAULT '' ",f.Name)
	case int:
		sql = fmt.Sprintf("\n `%s` integer NULL DEFAULT 0 ",f.Name)
	case int32:
		sql = fmt.Sprintf("\n `%s` integer NULL DEFAULT 0 ",f.Name)
	case int64:
		sql = fmt.Sprintf("\n `%s` integer NULL DEFAULT 0 ",f.Name)
	case float64:
		sql = fmt.Sprintf("\n `%s` float NULL DEFAULT 0.0 ",f.Name)
	case float32:
		sql = fmt.Sprintf("\n `%s` float NULL DEFAULT 0.0 ",f.Name)
	default:
		sql = fmt.Sprintf("\n `%s` varchar(255) NULL DEFAULT '' ",f.Name)
	}

	if f.Pk{
		sql = fmt.Sprintf("\n `%s` integer AUTO_INCREMENT PRIMARY KEY",f.Name)
	}

	sql += ","

	return sql
}

type DBModel struct {
	Name string
	Fields []Field
}

func (d *DBModel) TableSql() string{
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (",d.Name)
	for _,field := range d.Fields{
		sql += field.Sql()
	}
	sql = sql[:len(sql)-1]
	sql += "\n ) ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	return sql
}

func (d *DBModel) InsertSql() string{
	sql := fmt.Sprintf("INSERT `%s` SET ",d.Name)
	for i:= 1;i< len(d.Fields);i++{
		sql += fmt.Sprintf("`%s`=?,",d.Fields[i].Name)
	}
	sql = sql[:len(sql)-1]
	return sql
}

func (d *DBModel) InsertArgs() []interface{}{
	args := []interface{}{}
	for i:= 1;i< len(d.Fields);i++{
		rv := reflect.ValueOf(d.Fields[i].Value)
		switch rv.Kind(){
		case reflect.Array:
		case reflect.Slice:
			bytes,_ := json.Marshal(d.Fields[i].Value)
			args = append(args,string(bytes))
		default:
			args = append(args,rv.String())

		}
	}
	return args
}

func NewDBModel(name string,m map[string]interface{}) *DBModel{

	dbModel := &DBModel{Name:name,Fields:[]Field{}}
	dbModel.Fields = append(dbModel.Fields,Field{Name:strings.ToLower("FffId"),Pk:true,Value:1})
	for k,v := range m{
		dbModel.Fields = append(dbModel.Fields,Field{Name:strings.ToLower(k),Pk:false,Value:v})
	}

	return dbModel
}

