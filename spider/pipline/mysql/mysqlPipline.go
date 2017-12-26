package mysql

import (
	"sync"
)

type MysqlPipline struct {
	sync.Once
}

func NewMysqlPipline() *MysqlPipline {
	return &MysqlPipline{}
}

func (c *MysqlPipline) ProcessData(v []map[string]interface{}, taskName string, processName string) {
	for _,m :=range v{
		dbModel := NewDBModel(processName,m)

		CreateTable(dbModel)
		Add(dbModel)
	}
}
