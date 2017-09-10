package process

import (
	"YiSpider/common/model"
	"YiSpider/spider/process/rule"
)

type HtmlProcess struct {

}

func (h *HtmlProcess)Process(task *model.Task,htmlBytes []byte) interface{}{
	result := rule.TemplateProcess(task.Process.TemplateRule.Rule,htmlBytes)
	return result
}