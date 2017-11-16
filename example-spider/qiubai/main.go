package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
)

func main() {

	task := &model.Task{
		Id:   "qiiubai",
		Name: "qiubai",
		Request: []*model.Request{
			{
				Method: "get",
				Url:    "https://www.qiushibaike.com",
			},
		},
		Process: []model.Process{
			{
				RegUrl: []string{
					"/.*?/page/[0-9]+",
					"/hot/|/imgrank/|/text/|/history/|/pic/|/textnew/",
				},
				Type: "template",
				TemplateRule: model.TemplateRule{
					Rule: map[string]string{
						"node":        "array|.article",
						"url":         "attr.href|.contentHerf",
						"author":      "attr.alt|.author a img",
						"content":     "text|.content span",
						"like_num":    "text|.stats-vote i",
						"comment_num": "text|.stats-comments a i",
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
