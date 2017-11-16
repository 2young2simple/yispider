package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
)

func main() {

	task := &model.Task{
		Id:   "douban-movie",
		Name: "douban-movie",
		Request: []*model.Request{
			{
				Method:      "get",
				Url:         "https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-10000,20}",
				ProcessName: "movie",
			},
		},
		Process: []model.Process{
			{
				Name: "movie",
				Type: "json",
				JsonRule: model.JsonRule{
					Rule: map[string]string{
						"node":  "array|data",
						"rate":  "rate",
						"star":  "star",
						"id":    "id",
						"url":   "url",
						"title": "title",
						"cover": "cover",
						"casts": "casts",
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

/*
 douban-movie json

 {
    "id": "douban-movie",
    "Name": "douban-movie",
    "request": [
        {
            "url": "https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-10,20}",
            "method": "get",
            "type": "",
            "data": null,
            "header": null,
            "cookies": {
                "url": "",
                "data": ""
            },
            "process_name": "movie"
        }
    ],
    "process": [
        {
            "name": "movie",
            "reg_url": null,
            "type": "json",
            "template_rule": {
                "Rule": null
            },
            "json_rule": {
                "Rule": {
                    "casts": "casts",
                    "cover": "cover",
                    "id": "id",
                    "node": "array|data",
                    "rate": "rate",
                    "star": "star",
                    "title": "title",
                    "url": "url"
                }
            },
            "add_queue": null
        }
    ],
    "pipline": "file",
    "depth": 0,
    "end_count": 0
}

curl -d '{"id":"douban-movie","Name":"douban-movie","request":[{"url":"https://movie.douban.com/j/new_search_subjects?sort=T\u0026range=0,10\u0026tags=\u0026start={0-100,20}","method":"get","type":"","data":null,"header":null,"cookies":{"url":"","data":""},"process_name":"movie"}],"process":[{"name":"movie","reg_url":null,"type":"json","template_rule":{"Rule":null},"json_rule":{"Rule":{"casts":"casts","cover":"cover","id":"id","node":"array|data","rate":"rate","star":"star","title":"title","url":"url"}},"add_queue":null}],"pipline":"file","depth":0,"end_count":0}' "http://127.0.0.1:7774/task/addAndRun"


*/
