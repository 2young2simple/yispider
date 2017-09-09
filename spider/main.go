package main

import (
	"YiSpider/spider/config"
	"YiSpider/spider/logger"
	"YiSpider/spider/downloader"
)

func main(){

	var err error

	if err = config.InitConfig();err != nil{
		logger.Info(err.Error())
		return
	}

	downloader.InitDownloader(config.ConfigI.WorkNum)

}
