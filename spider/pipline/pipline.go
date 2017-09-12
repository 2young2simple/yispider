package pipline

import "YiSpider/spider/model"

type Pipline interface {
	ProcessData(v interface{},task *model.Task)
}

