package core

import (
	"YiSpider/spider/schedule"
	"YiSpider/spider/config"
	"YiSpider/spider/logger"
	"YiSpider/spider/downloader"
	"YiSpider/spider/model"
	"time"
	"sync"
	"sync/atomic"
	"net/http"
	"YiSpider/spider/spider"
)

const Default_WorkNum  = 1

type SpiderRuntime struct {
	sync.Mutex
	workNum  int
	schedule *schedule.Schedule
	spider   *spider.Spider

	stopSign  bool
	recoverChan  chan int

	TaskMeta *TaskMeta
}

type TaskMeta struct {
	DownloadFailCount int32 `json:"download_fail_count"`
	DownloadCount int32 `json:"download_fail_count"`

	UrlNum int32 `json:"url_num"`
	WaitUrlNum int `json:"wait_url_num"`
	CrawlerResultNum int32 `json:"crawler_result_num"`
}

func NewSpiderRuntime() *SpiderRuntime{

	workNum := config.ConfigI.WorkNum
	if workNum == 0{
		workNum = Default_WorkNum
	}

	s := &SpiderRuntime{}
	s.workNum = workNum
	s.schedule = schedule.NewSchedule(config.ConfigI.MaxWaitNum)
	s.recoverChan = make(chan int)
	meta := &TaskMeta{}
	meta.WaitUrlNum = 0
	meta.UrlNum = int32(0)
	meta.DownloadCount = int32(0)
	meta.DownloadFailCount = int32(0)
	meta.CrawlerResultNum = int32(0)

	s.TaskMeta = meta

	return s
}


func (s *SpiderRuntime)SetSpider(spider *spider.Spider) {
	s.spider = spider
}

func (s *SpiderRuntime)GetSpider() *spider.Spider{
	return s.spider
}


func (s *SpiderRuntime)Run(){
	if s.stopSign{
		s.recoverChan <- 1
		return
	}
	s.schedule.PushMuti(s.spider.GetRequests())

	for i:=0;i<s.workNum;i++{
		go s.worker()
	}
}

func (s *SpiderRuntime)Stop(){
	s.stopSign = true
}


func (s *SpiderRuntime) worker(){

	for{
		if s.stopSign{
			_,ok := <- s.recoverChan
			s.stopSign = false
			if !ok{
				goto exit
			}
		}

		req,ok := s.schedule.Pop()
		if !ok{
			goto exit
		}

		atomic.AddInt32(&s.TaskMeta.DownloadCount,1)
		response,err := s.download(req)
		if err != nil{
			logger.Error(err.Error())
			atomic.AddInt32(&s.TaskMeta.DownloadFailCount,1)
			continue
		}
		p,ok := s.spider.Process[req.ProcessName]
		if !ok{
			logger.Info("process is not find ! please call SetProcess|SetTask")
			break
		}

		page,err := p.Process(response)
		if err!= nil{
			logger.Info("Process error|",err.Error())
			break
		}

		atomic.AddInt32(&s.TaskMeta.UrlNum,int32(len(page.Urls)))

		s.TaskMeta.WaitUrlNum = s.schedule.Count()

		go func (){
			s.schedule.PushMuti(page.Urls)
		}()

		atomic.AddInt32(&s.TaskMeta.CrawlerResultNum,int32(page.ResultCount))

		s.spider.Pipline.ProcessData(page.Result,s.spider.Name,req.ProcessName)
	}

exit:
	logger.Info(s.spider.Name,"worker close")
}


func (s *SpiderRuntime) download(req *model.Request) (*http.Response,error){
	time.Sleep(1*time.Second)
	switch req.Method {
	case "get":
		return downloader.Get(req.ProcessName,req.Url)
	case "post":
		return downloader.PostJson(req.ProcessName,req.Url,req.Data)
	}

	return nil,nil
}

func (s *SpiderRuntime) Exit(){
	s.schedule.Close()
	close(s.recoverChan)
}

