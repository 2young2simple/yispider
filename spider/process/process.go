package process

import "YiSpider/common/model"

type Process interface {
	Process(bytes []byte,task *model.Task) interface{}
}