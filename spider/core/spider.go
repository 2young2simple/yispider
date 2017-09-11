package core

import (
	"YiSpider/spider/schedule"
	"YiSpider/spider/config"
	"YiSpider/spider/logger"
	"YiSpider/spider/pipline"
	"YiSpider/spider/process"
	"YiSpider/spider/downloader"
	"YiSpider/common/model"
	"YiSpider/spider/pipline/file"
	"YiSpider/spider/pipline/console"
)

const Default_WorkNum  = 4

type Spider struct {
	workNum  int
	piplines map[string]pipline.Pipline
	processMap map[string]process.Process
	schedule *schedule.Schedule
	closeChan  chan int
}

func NewSpider() *Spider{

	var err error
	if err = config.InitConfig();err != nil{
		logger.Info(err.Error())
		return nil
	}

	workNum := config.ConfigI.MaxWaitNum
	if workNum == 0{
		workNum = Default_WorkNum
	}
	sche := schedule.NewSchedule(config.ConfigI.MaxWaitNum)

	spider := &Spider{}
	spider.workNum = workNum
	spider.schedule = sche
	spider.piplines = make(map[string]pipline.Pipline)
	spider.processMap = make(map[string]process.Process)
	spider.closeChan = make(chan int)

	spider.registerDefaultProcess()
	spider.registerDefaultPipline()

	return spider
}

func (s *Spider)registerDefaultProcess(){
	s.AddProcess("html",process.NewHtmlProcess())
}
func (s *Spider)registerDefaultPipline() {
	s.AddPipline("file",file.NewFilePipline("./"))
	s.AddPipline("console",console.NewConsolePipline())
}

func (s *Spider)AddProcess(key string,process process.Process) *Spider{
	s.processMap[key] = process
	return s
}

func (s *Spider)AddTask(task *model.Task){
	s.schedule.Push(task)
}

func (s *Spider)AddPipline(key string,pipline pipline.Pipline) *Spider{
	s.piplines[key] = pipline
	return s
}

func (s *Spider)Run(){

	for i:=0;i<s.workNum;i++{
		go s.worker()
	}
	select {}
}



func (s *Spider) worker(){

	for{
		task := s.schedule.Pop()
		bytes,err := s.download(task)
		if err != nil{
			logger.Info("Download fail, task:",task.Name,"url:",task.Url)
			continue
		}

		curProcess := s.getPageProcess(task)
		if curProcess == nil{
			logger.Info("getPageProcess fail, not find the process key:",task.Process.Type)
			continue
		}

		result := curProcess.Process(bytes,task)

		pip,ok := s.piplines[task.Pipline]
		if !ok{
			logger.Info("get Pipline fail, not find the pipline key:",task.Pipline)
			continue
		}
		pip.ProcessData(result,task)

	}
}

func (s *Spider) getPageProcess(task *model.Task) process.Process{
	switch task.Process.Type{
	case "html":
		return s.processMap["html"]
	case "json":
		return s.processMap["json"]
	}
	return nil
}

func (s *Spider) download(task *model.Task) ([]byte,error){
	switch task.Method {
	case "get":
		return downloader.Get(task.Id,task.Url)
	case "post":
		return downloader.PostJson(task.Id,task.Url,task.RequestBody.Data)
	}

	return []byte{},nil
}

func (s *Spider) Exit(){

}

