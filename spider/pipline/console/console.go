package console

import (
	"fmt"
	"YiSpider/common/model"
	"encoding/json"
)

type ConsolePipline struct{

}

func NewConsolePipline() *ConsolePipline{
	return &ConsolePipline{}
}

func (c *ConsolePipline)ProcessData(v interface{},task *model.Task){
	bytes,_ := json.Marshal(v)
	fmt.Println("Pipline :",string(bytes))
}