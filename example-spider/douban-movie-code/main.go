package main

import (
	"YiSpider/spider/model"
	"YiSpider/spider"
)

func main(){

	task := &model.Task{
		Id:"douban-movie",
		Name:"douban-movie",
		Method:"get",
		Url:"https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start=0",
		Process: model.Process{
			Url:"https://movie.douban.com/j/new_search_subjects?sort=T&range=0,100&tags=&start=0",
			RegUrl:[]string{},
			Type:"json",
			JsonRule:model.JsonRule{
				Rule:map[string]string{
					"node":"array|data",
					"rate":"rate",
					"star":"star",
					"id":"id",
					"url":"url",
					"title":"title",
				},
			},
		},
		Pipline:"file",
	}

	app := spider.New()
	app.AddTask(task)
	app.Run()
}
