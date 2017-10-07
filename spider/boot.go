package main

import (
	"YiSpider/spider/core"
	"YiSpider/spider/model"
	"YiSpider/spider/http"
	"YiSpider/spider/register/etcd"
	"YiSpider/spider/config"
	"YiSpider/spider/logger"
)

func main(){

	var err error
	if err = config.InitConfig();err != nil{
		logger.Info(err.Error())
	}

	task := &model.Task{
		Id:"qiiubai",
		Name:"qiubai",
		Method:"get",
		Host:"https://www.qiushibaike.com",
		Url:"https://www.qiushibaike.com",
		Process: model.Process{
			Url:"https://www.qiushibaike.com",
			RegUrl:[]string{
				"/.*?/page/[0-9]+",
				"/hot/|/imgrank/|/text/|/history/|/pic/|/textnew/",
			},
			Type:"template",
			TemplateRule:model.TemplateRule{
				Rule:map[string]string{
					"node":"array|.article",
					"url":"attr.href|.contentHerf",
					"author":"attr.alt|.author a img",
					"content":"text|.content span",
					"like_num":"text|.stats-vote i",
					"comment_num":"text|.stats-comments a i",
				},
			},
		},
		Pipline:"file",
	}

	task1 := &model.Task{
		Id:"sohu",
		Name:"sohu",
		Method:"get",
		Host:"https://www.qiushibaike.com",
		Url:"https://www.qiushibaike.com",
		Process: model.Process{
			Url:"https://www.qiushibaike.com",
			RegUrl:[]string{
				"/.*?/page/[0-9]+",
				"/hot/|/imgrank/|/text/|/history/|/pic/|/textnew/",
			},
			Type:"template",
			TemplateRule:model.TemplateRule{
				Rule:map[string]string{
					"node":"array|.article",
					"url":"attr.href|.contentHerf",
					"author":"attr.alt|.author a img",
					"content":"text|.content span",
					"like_num":"text|.stats-vote i",
					"comment_num":"text|.stats-comments a i",
				},
			},
		},
		Pipline:"file",
	}


	app := core.New()
	app.AddTask(task)
	app.AddTask(task1)
	app.Run()

	worker := etcd.NewWorker(config.ConfigI.Name,config.ConfigI.HttpAddr,config.ConfigI.Etcd)
	go worker.HeartBeat()


	http.InitHttpServer()

}
