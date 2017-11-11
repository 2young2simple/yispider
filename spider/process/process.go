package process

import (
	"YiSpider/spider/model"
)


type Process interface {
	Process(context model.Context) (*model.Page,error)
}
