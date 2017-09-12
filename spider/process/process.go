package process

import "YiSpider/spider/model"

type Process interface {
	Process(bytes []byte,task *model.Task) *model.Page
}