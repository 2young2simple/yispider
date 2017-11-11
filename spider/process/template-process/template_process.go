package template_process

import (
	"YiSpider/spider/model"
)

type TemplateProcess struct {
	tempProcess *model.Process
}

func NewTemplateProcess(tempProcess *model.Process) *TemplateProcess{
		return &TemplateProcess{tempProcess: tempProcess}
}

func (t *TemplateProcess)Process(context model.Context) (*model.Page,error){
	return TemplateRuleProcess(t.tempProcess,context)

}



