package main

import (
	"YiSpider/spider/core"
	"YiSpider/common/model"
)

func main(){
	spider := core.NewSpider()
	task := &model.Task{
		Id:"qiiubai",
		Name:"qiubai",
		Method:"get",
		Url:"https://www.qiushibaike.com",
		Process: model.Process{
			Url:"https://www.qiushibaike.com",
			Type:"html",
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
	spider.AddTask(task)

	spider.Run()
}
