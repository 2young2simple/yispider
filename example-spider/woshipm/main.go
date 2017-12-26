package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
)

func main() {

	task := &model.Task{
		Id:   "woshipm",
		Name: "woshipm",
		Request: []*model.Request{
			{
				Method:      "get",
				Url:         "http://www.woshipm.com/category/pd/page/{1-588,1}",
				ProcessName: "woshipm-list",
			},
		},
		Process: []model.Process{
			{
				Name: "woshipm-list",
				Type: "template",
				TemplateRule: model.TemplateRule{
					Rule: map[string]string{
						"node":     "array|.postlist-item",
						"img":      "attr.src|.post-img a img",
						"time":     "text|.stream-list-meta time",
						"title":    "text|.post-title a",
						"author":   "text|.author a",
						"des": "text|.des",
						"read_num":   "text|.post-meta-items span:nth-child(1)",
						"collect_num":   "text|.post-meta-items span:nth-child(2)",
						"like_num":   "text|.post-meta-items span:nth-child(3)",
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
