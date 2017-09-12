package process

import (
	"YiSpider/common/model"
	"YiSpider/spider/process/rule"
)

type HtmlProcess struct {

}

func NewHtmlProcess() *HtmlProcess{
	return &HtmlProcess{}
}

func (h *HtmlProcess)Process(bytes []byte,task *model.Task) interface{}{
	result := rule.TemplateProcess(task.Process.TemplateRule.Rule,bytes)
	return result
}