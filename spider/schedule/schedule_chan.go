package schedule

import (
	"YiSpider/spider/model"
	"YiSpider/manage/logger"
)

type Schedule struct{
	waitQueue chan *model.Task
}

func NewSchedule(maxWaitNum int) *Schedule{
	schedule := &Schedule{}
	schedule.waitQueue = make(chan *model.Task,maxWaitNum)
	return schedule
}


func (d *Schedule) Push(task *model.Task){
	logger.Info("Push Url:",task.Process.Url)
	d.waitQueue <- task
}

func (d *Schedule) Pop() (*model.Task,bool){
	 task,ok := <- d.waitQueue
	return task,ok
}

func (d *Schedule)Close() {
	close(d.waitQueue)
}


