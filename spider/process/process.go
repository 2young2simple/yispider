package process

import (
	"YiSpider/spider/model"
	"net/http"
)


type Process interface {
	Process(response *http.Response) (*model.Page,error)
}
