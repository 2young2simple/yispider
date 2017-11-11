package json_process

import (
	"YiSpider/spider/model"
)

type JsonProcess struct {
	jsonProcess *model.Process

}

func NewJsonProcess(jsonProcess *model.Process) *JsonProcess{
		return &JsonProcess{jsonProcess:jsonProcess}
}

func (j *JsonProcess)Process(context model.Context) (*model.Page,error){
	return JsonRuleProcess(j.jsonProcess,context)
}