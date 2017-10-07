package core

import (
	"YiSpider/spider/schedule"
	"YiSpider/spider/config"
	"YiSpider/spider/logger"
	"YiSpider/spider/pipline"
	"YiSpider/spider/process"
	"YiSpider/spider/downloader"
	"YiSpider/spider/model"
	"YiSpider/spider/pipline/file"
	"YiSpider/spider/pipline/console"
	"encoding/json"
	"time"
	"sync"
)

const Default_WorkNum  = 4

type SpiderRuntime struct {
	sync.Mutex
	workNum  int
	piplines map[string]pipline.Pipline
	processMap map[string]process.Process
	schedule *schedule.Schedule
	task   *model.Task

	stopSign  bool
	recoverChan  chan int
}

func NewSpiderRuntime() *SpiderRuntime{

	workNum := config.ConfigI.WorkNum
	if workNum == 0{
		workNum = Default_WorkNum
	}

	spider := &SpiderRuntime{}
	spider.workNum = workNum
	spider.schedule = schedule.NewSchedule(config.ConfigI.MaxWaitNum)
	spider.piplines = make(map[string]pipline.Pipline)
	spider.processMap = make(map[string]process.Process)
	spider.recoverChan = make(chan int)

	spider.registerDefaultProcess()
	spider.registerDefaultPipline()

	return spider
}

func (s *SpiderRuntime)registerDefaultProcess(){
	s.AddProcess("template",process.NewTemplateProcess())
}

func (s *SpiderRuntime)registerDefaultPipline() {
	s.AddPipline("file",file.NewFilePipline("./"))
	s.AddPipline("console",console.NewConsolePipline())
}

func (s *SpiderRuntime)AddProcess(key string,process process.Process) *SpiderRuntime{
	s.processMap[key] = process
	return s
}

func (s *SpiderRuntime)SetTask(task *model.Task){
	s.task = task
}

func (s *SpiderRuntime)GetTask() *model.Task{
	return s.task
}


func (s *SpiderRuntime)AddPipline(key string,pipline pipline.Pipline) *SpiderRuntime{
	s.piplines[key] = pipline
	return s
}

func (s *SpiderRuntime)Run(){

	s.schedule.Push(s.task)

	for i:=0;i<s.workNum;i++{
		go s.worker()
	}
}

func (s *SpiderRuntime)Stop(){
	s.recoverChan <- 1
}


func (s *SpiderRuntime) worker(){

	for{
		if s.stopSign{
			_,ok := <- s.recoverChan
			if !ok{
				goto exit
			}
		}

		task,ok := s.schedule.Pop()
		if !ok{
			goto exit
		}
		logger.Info("Pop Url:",task.Process.Url)

		bytes,err := s.download(task)
		if err != nil{
			continue
		}

		curProcess := s.getPageProcess(task)
		if curProcess == nil{
			logger.Info("getPageProcess fail, not find the process key:",task.Process.Type)
			continue
		}

		page := curProcess.Process(bytes,task)

		for _,url := range page.Urls{
			logger.Info("Dicover Url:",task.Process.Url)

				//TODO BIG BIG BUG ... deadlock repeat
			aj, _ := json.Marshal(task)
			copy := new(model.Task)
			_ = json.Unmarshal(aj,copy)
			copy.Process.Url = task.Host+url
			s.schedule.Push(copy)
		}

		pip,ok := s.piplines[task.Pipline]
		if !ok{
			logger.Info("get Pipline fail, not find the pipline key:",task.Pipline)
			continue
		}
		pip.ProcessData(page.Result,task)
	}

exit:
	logger.Info(s.task.Name,"worker close")
}

func (s *SpiderRuntime) getPageProcess(task *model.Task) process.Process{
	switch task.Process.Type{
	case "template":
		return s.processMap["template"]
	case "json":
		return s.processMap["json"]
	}
	return nil
}

func (s *SpiderRuntime) download(task *model.Task) ([]byte,error){
	time.Sleep(1*time.Second)
	switch task.Method {
	case "get":
		//logger.Info("Download Url :",task.Process.Url)
		return downloader.Get(task.Id,task.Process.Url)
	case "post":
		return downloader.PostJson(task.Id,task.Process.Url,task.RequestBody.Data)
	}

	return []byte{},nil
}

func (s *SpiderRuntime) Exit(){
	s.schedule.Close()
	close(s.recoverChan)
}

