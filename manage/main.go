package main

import (
	"YiSpider/manage/config"
	"YiSpider/manage/discover"
	"YiSpider/manage/http"
	"YiSpider/manage/logger"
)

func main() {

	var err error

	if err = config.InitConfig(); err != nil {
		logger.Info(err.Error())
		return
	}

	discover.InitDiscover()

	http.InitHttpServer()

}
