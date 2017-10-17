package schedule

import (
	"YiSpider/spider/model"
	"YiSpider/manage/logger"
	"YiSpider/spider/common"
)

type Schedule struct{
	waitQueue chan *model.Request
}

func NewSchedule(maxWaitNum int) *Schedule{
	schedule := &Schedule{}
	schedule.waitQueue = make(chan *model.Request,maxWaitNum)
	return schedule
}


func (d *Schedule) Push(req *model.Request){
	praseReqs := common.PraseReq([]*model.Request{req},nil)
	for _,req := range praseReqs{
		logger.Info("Push Url:",req.Url,req.ProcessName,len(d.waitQueue))
		d.waitQueue <- req
	}
}

func (d *Schedule) PushMuti(reqs []*model.Request){
	praseReqs := common.PraseReq(reqs,nil)
	for _,req := range praseReqs{
		logger.Info("Push Url:",req.Url,req.ProcessName,len(d.waitQueue))
		d.waitQueue <- req
	}
}

func (d *Schedule) Pop() (*model.Request,bool){
	 req,ok := <- d.waitQueue
	return req,ok
}

func (d *Schedule) Count() int{
	return len(d.waitQueue)
}

func (d *Schedule)Close() {
	close(d.waitQueue)
}


