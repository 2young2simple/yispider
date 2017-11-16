package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
)

func main() {

	task := &model.Task{
		Id:   "qiongyou",
		Name: "qiongyou",
		Request: []*model.Request{
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_1_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_2_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_3_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_4_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_5_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_6_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_7_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_8_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_9_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_10_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
			{
				Method:      "get",
				Url:         "http://plan.qyer.com/search_0_0_0_0_0_11_{1-134,1}/",
				ProcessName: "qiongyou-list",
			},
		},
		Process: []model.Process{
			{
				Name: "qiongyou-list",
				Type: "template",
				TemplateRule: model.TemplateRule{
					Rule: map[string]string{
						"node":     "array|.items",
						"img":      "attr.src|.plan-cover",
						"time":     "text|.fontYaHei dt",
						"title":    "text|.fontYaHei dd",
						"day":      "text|.day strong",
						"tag":      "text|.tag strong",
						"plan":     "text|.plan p",
						"author":   "text|.name",
						"read_num": "text|.number .icon1",
						"xx_num":   "text|.number .icon2",
					},
				},
			},
		},

		Pipline: "file",
	}

	app := spider.New()
	app.AddSpider(spider2.InitWithTask(task))
	app.Run()
}
