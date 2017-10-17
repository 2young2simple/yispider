package json_process

import (
	"YiSpider/spider/model"
	"net/http"
	"io/ioutil"
)

type JsonProcess struct {
	jsonProcess *model.Process

}

func NewJsonProcess(jsonProcess *model.Process) *JsonProcess{
		return &JsonProcess{jsonProcess:jsonProcess}
}

func (j *JsonProcess)Process(response *http.Response) (*model.Page,error){
	body,err := ioutil.ReadAll(response.Body)
	if err != nil{
		return nil,err
	}
	defer response.Body.Close()
	return JsonRuleProcess(j.jsonProcess,body)
}