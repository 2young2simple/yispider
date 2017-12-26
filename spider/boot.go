package spider

import (
	"YiSpider/spider/config"
	"YiSpider/spider/core"
	"YiSpider/spider/http"
	"YiSpider/spider/register/etcd"
	"YiSpider/spider/spider"
)

type Boot struct {
	engine *core.Engine
}

func init() {
	var err error

	if err = config.InitConfig(); err != nil {
		panic(err)
	}
}

func New() *Boot {
	s := &Boot{}
	s.engine = core.New()
	return s
}

func (s *Boot) AddSpider(spider *spider.Spider) *core.Engine {
	return s.engine.AddSpider(spider)
}

func (s *Boot) Run() {

	s.engine.Run()

	if len(config.ConfigI.Etcd) > 0{
		worker := etcd.NewWorker(config.ConfigI.Name, config.ConfigI.HttpAddr, config.ConfigI.Etcd)
		go worker.HeartBeat()
	}

	http.InitHttpServer()
}
