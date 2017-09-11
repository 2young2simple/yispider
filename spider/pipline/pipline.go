package pipline

import "YiSpider/common/model"

type Pipline interface {
	ProcessData(v interface{},task *model.Task)
}

