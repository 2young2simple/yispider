package schedule

import (
	"YiSpider/common/model"
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
	d.waitQueue <- task
}

func (d *Schedule) Pop() *model.Task{
	return <- d.waitQueue
}


