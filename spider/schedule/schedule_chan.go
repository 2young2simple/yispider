package schedule

import (
	"YiSpider/manage/logger"
	"YiSpider/spider/common"
	"YiSpider/spider/config"
	"YiSpider/spider/model"
)

type ChanSchedule struct {
	waitQueue chan *model.Request
}

func NewChanSchedule(config *config.Config) Schedule {
	schedule := &ChanSchedule{}
	schedule.waitQueue = make(chan *model.Request, config.MaxWaitNum)
	return schedule
}

func (d *ChanSchedule) Push(req *model.Request) {
	praseReqs := common.PraseReq([]*model.Request{req}, nil)
	for _, req := range praseReqs {
		logger.Info("Push Url:", req.Url, req.ProcessName, len(d.waitQueue))
		d.waitQueue <- req
	}
}

func (d *ChanSchedule) PushMuti(reqs []*model.Request) {
	praseReqs := common.PraseReq(reqs, nil)
	for _, req := range praseReqs {
		logger.Info("Push Url:", req.Url, req.ProcessName, len(d.waitQueue))
		d.waitQueue <- req
	}
}

func (d *ChanSchedule) Pop() (*model.Request, bool) {
	req, ok := <-d.waitQueue
	logger.Info("Pop Url:", req.Url, req.ProcessName, len(d.waitQueue))
	return req, ok
}

func (d *ChanSchedule) Count() int {
	return len(d.waitQueue)
}

func (d *ChanSchedule) Close() {
	close(d.waitQueue)
}

func init() {
	RegisterSchedule("chan", NewChanSchedule)
}
