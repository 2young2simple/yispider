package template_process

import (
	"YiSpider/spider/model"
	"net/http"
)

type TemplateProcess struct {
	tempProcess *model.Process
}

func NewTemplateProcess(tempProcess *model.Process) *TemplateProcess{
		return &TemplateProcess{tempProcess: tempProcess}
}

func (t *TemplateProcess)Process(response *http.Response) (*model.Page,error){
	return TemplateRuleProcess(t.tempProcess,response)
}



