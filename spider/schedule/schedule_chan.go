package schedule

import (
	"YiSpider/common/model"
	"YiSpider/spider/downloader"
)

type Schedule struct{
	taskQ	  chan *model.Task
	closeChan chan int
}

var ScheduleI *Schedule


func InitDownloader(workNum int){
	ScheduleI = &Schedule{}
	ScheduleI.taskQ = make(chan *model.Task,workNum*2)
	ScheduleI.closeChan = make(chan int)

	for i:=0;i<workNum;i++{
		go ScheduleI.worker()
	}
}

func (d *Schedule) worker(){
	for{
		select{
		case task := <- d.taskQ:
			d.download(task)
		case <- d.closeChan:
			break
		}

	}
}

func (d *Schedule) Push(task *model.Task){
	d.taskQ <- task
}

func (d *Schedule) download(task *model.Task){
	switch task.Method {
	case "get":
		downloader.Get(task.Id,task.Url)
	case "post":
		downloader.PostJson(task.Id,task.Url,task.RequestBody.Data)
	}
}

func (d *Schedule) Exit(){
	d.closeChan <- 1
}

