package core

import (
	"YiSpider/spider/config"
	"YiSpider/spider/downloader"
	"YiSpider/spider/logger"
	"YiSpider/spider/model"
	"YiSpider/spider/schedule"
	//"time"
	"YiSpider/spider/common"
	"YiSpider/spider/process"
	"YiSpider/spider/spider"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
)

const Default_WorkNum = 1

type SpiderRuntime struct {
	sync.Mutex
	workNum  int
	schedule schedule.Schedule
	spider   *spider.Spider

	stopSign    bool
	recoverChan chan int

	TaskMeta *TaskMeta
}

type TaskMeta struct {
	DownloadFailCount int32 `json:"download_fail_count"`
	DownloadCount     int32 `json:"download_fail_count"`

	UrlNum           int32 `json:"url_num"`
	WaitUrlNum       int   `json:"wait_url_num"`
	CrawlerResultNum int32 `json:"crawler_result_num"`
}

func NewSpiderRuntime() *SpiderRuntime {

	workNum := config.ConfigI.WorkNum
	if workNum == 0 {
		workNum = Default_WorkNum
	}

	s := &SpiderRuntime{}
	s.workNum = workNum
	s.schedule = schedule.GetSchedule(config.ConfigI)
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

func (s *SpiderRuntime) SetSpider(spider *spider.Spider) {
	s.spider = spider
}

func (s *SpiderRuntime) GetSpider() *spider.Spider {
	return s.spider
}

func (s *SpiderRuntime) Run() {
	if s.stopSign {
		s.recoverChan <- 1
		return
	}
	s.schedule.PushMuti(s.spider.GetRequests())

	for i := 0; i < s.workNum; i++ {
		go s.worker()
	}
}

func (s *SpiderRuntime) Stop() {
	s.stopSign = true
}

func (s *SpiderRuntime) worker() {
	context := model.Context{}

	for {
		if s.stopSign {
			_, ok := <-s.recoverChan
			s.stopSign = false
			if !ok {
				goto exit
			}
		}

		req, ok := s.schedule.Pop()
		if !ok {
			goto exit
		}
		if req == nil {
			logger.Info("schedule is emply")
			continue
		}

		atomic.AddInt32(&s.TaskMeta.DownloadCount, 1)
		response, err := s.download(req)
		if err != nil {
			logger.Error(err.Error())
			atomic.AddInt32(&s.TaskMeta.DownloadFailCount, 1)
			continue
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		context.Clear()
		context.Body, err = common.ToUtf8(body)
		if err != nil {
			context.Body = body
		}
		context.Request = response.Request
		context.Header = response.Header

		ps, ok := s.spider.Process[req.ProcessName]
		if !ok {
			response.Body.Close()
			logger.Info("process is not find ! please call SetProcess|SetTask")
			break
		}
		for _, p := range ps {
			page, err := processWrapper(p, context)
			if err != nil {
				logger.Info("Process fail|", err.Error())
				continue
			}
			if page == nil {
				logger.Info("Process page is nil")
				continue
			}
			s.TaskMeta.WaitUrlNum = s.schedule.Count()

			if page.Urls != nil && len(page.Urls) > 0 {
				atomic.AddInt32(&s.TaskMeta.UrlNum, int32(len(page.Urls)))
				go func() {
					s.schedule.PushMuti(page.Urls)
				}()
			}

			if page.ResultCount > 0 {

				atomic.AddInt32(&s.TaskMeta.CrawlerResultNum, int32(page.ResultCount))

				s.spider.Pipline.ProcessData(page.Result, s.spider.Name, req.ProcessName)
			}
		}

		response.Body.Close()
	}

exit:
	logger.Info(s.spider.Name, "worker close")
}
func processWrapper(p process.Process, context model.Context) (*model.Page, error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()

	page, err := p.Process(context)
	return page, err
}

func (s *SpiderRuntime) download(req *model.Request) (*http.Response, error) {
	//time.Sleep(1*time.Second)
	switch req.Method {
	case "get":
		return downloader.Get(req.ProcessName, req.Url)
	case "post":
		return downloader.PostJson(req.ProcessName, req.Url, req.Data)
	}

	return nil, nil
}

func (s *SpiderRuntime) Exit() {
	s.schedule.Close()
	close(s.recoverChan)
}
