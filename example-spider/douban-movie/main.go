package qiubai

import (
	"YiSpider/spider/model"
	"YiSpider/spider"
)

func main(){

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

	app := spider.New()
	app.AddTask(task)
	app.AddTask(task1)
	app.Run()
}
