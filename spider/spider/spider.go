package template_spider

import (
	"YiSpider/spider/model"
	"YiSpider/spider/process"
	"YiSpider/spider/pipline"
)

type Spider struct {
	Id   string
	Name string

	Depth int
	EndCount int

	Requests []model.Request

	Process map[string]process.Process
	Pipline pipline.Pipline
}


func (s *Spider)GetPipline() pipline.Pipline{
	return s.Pipline
}

func (s *Spider)GetProcess(name string) process.Process{
	return s.Process[name]
}

func (s *Spider)GetRequests() []model.Request{
	return s.Requests
}

