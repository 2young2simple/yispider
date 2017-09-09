package main

import (
	"YiSpider/manage/logger"
	"YiSpider/manage/config"
)

func main(){

	var err error

	if err = config.InitConfig();err != nil{
		logger.Info(err.Error())
		return
	}

}
