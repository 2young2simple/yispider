package main

import (
	"YiSpider/manage/logger"
	"YiSpider/manage/config"
	"YiSpider/manage/discover"
	"YiSpider/manage/http"
)

func main(){

	var err error

	if err = config.InitConfig();err != nil{
		logger.Info(err.Error())
		return
	}

	discover.InitDiscover()

	http.InitHttpServer()

}
