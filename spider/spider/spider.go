package spider

import (
	"YiSpider/spider/model"
	"YiSpider/spider/process"
	"YiSpider/spider/pipline"
	"YiSpider/spider/process/template-process"
	"YiSpider/spider/process/json-process"
	"YiSpider/spider/pipline/console"
	"YiSpider/spider/pipline/file"
)

type Spider struct {
	Id   string
	Name string

	Depth int
	EndCount int

	Requests []*model.Request

	Process map[string][]process.Process
	Pipline pipline.Pipline
}


func (s *Spider)GetPipline() pipline.Pipline{
	return s.Pipline
}

func (s *Spider)GetProcess(name string) []process.Process{
	return s.Process[name]
}

func (s *Spider)GetRequests() []*model.Request{
	return s.Requests
}

func InitWithTask(task *model.Task) *Spider {
	s := &Spider{}
	s.Id = task.Id
	s.Name = task.Name
	s.Depth = task.Depth
	s.EndCount = task.EndCount
	s.Requests = task.Request

	s.Process = make(map[string][]process.Process)

	for i,p := range task.Process{
		switch p.Type {
		case "template":
			processs,ok := s.Process[p.Name]
			if !ok {
				processs = []process.Process{}
				s.Process[p.Name] = processs
			}
			s.Process[p.Name] = append(processs,template_process.NewTemplateProcess(&task.Process[i]))
		case "json":
			processs,ok := s.Process[p.Name]
			if !ok {
				processs = []process.Process{}
				s.Process[p.Name] = processs
			}
			s.Process[p.Name] = append(processs,json_process.NewJsonProcess(&task.Process[i]))
		}
	}

	switch task.Pipline {
	case "console":
		s.Pipline = console.NewConsolePipline()
	case "file":
		s.Pipline = file.NewFilePipline("./")
	}

	return s
}


