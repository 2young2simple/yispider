package main

import (
	"YiSpider/spider/model"
	"YiSpider/spider"
	spider2 "YiSpider/spider/spider"
)

func main(){

	task := &model.Task{
		Id:"tuiku",
		Name:"tuiku",
		Request:[]*model.Request{
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/0/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/101000000/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/101040000/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/101050000/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/20/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/108000000/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
			{
				Method:"get",
				Url:"http://www.tuicool.com/ah/114000000/{1-100,1}?lang=1",
				ProcessName:"tuikulist",
			},
		},
		Process: []model.Process{
			{
				Name:"tuikulist",
				Type:"template",
				TemplateRule:model.TemplateRule{
					Rule:map[string]string{
						"node":"array|.list_article_item",
						"img":"attr.src|.article_thumb_image img",
						"title":"text|.title a",
						"author":"text|.tip span:nth-child(1)",
						"time":"text|.tip span:nth-child(3)",
					},
				},
			},
		},

		Pipline:"file",
	}

	app := spider.New()
	app.AddSpider(spider2.InitWithTask(task))
	app.Run()
}
