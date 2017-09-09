package downloader

import (
	"YiSpider/common/model"
)

type Downloader struct{
	taskQ	  chan *model.Task
	closeChan chan int
}

var DownloaderI *Downloader


func InitDownloader(workNum int){
	DownloaderI = &Downloader{}
	DownloaderI.taskQ = make(chan *model.Task,workNum*2)
	DownloaderI.closeChan = make(chan int)

	for i:=0;i<workNum;i++{
		go DownloaderI.worker()
	}
}

func (d *Downloader) worker(){
	for{
		select{
		case task := <- d.taskQ:
			d.download(task)
		case <- d.closeChan:
			break
		}

	}
}

func (d *Downloader) Push(task *model.Task){
	d.taskQ <- task
}

func (d *Downloader) download(task *model.Task){
	switch task.Method {
	case "get":
		Get(task.Id,task.Url)
	case "post":
		PostJson(task.Id,task.Url,task.RequestBody.Data)
	}
}

func (d *Downloader) Exit(){
	d.closeChan <- 1
}

