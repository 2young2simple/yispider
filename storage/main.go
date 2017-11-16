package main

import (
	"YiSpider/storage/config"
	"YiSpider/storage/logger"
)

func main() {

	var err error

	if err = config.InitConfig(); err != nil {
		logger.Info(err.Error())
		return
	}

}
