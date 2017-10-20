# YiSpider
A distributed spider platform

## 介绍

一款分布式爬虫平台，帮助你更好的管理和开发爬虫。
内置一套爬虫定义规则（模版），可使用模版快速定义爬虫，也可当作框架手动开发爬虫

## 架构

目前框架分为2个部分:  
#### 1.爬虫部分（spider节点）:

内部结构参考python scrapy框架，主要由 schedule,page process,pipline 4个部分组成，单个爬虫单独调度器，单独上下文管理,目前内置2中pipline的方式，控制台和文件,节点信息注册在etcd上用于manage节点发现。 

* `core`:负责爬虫生命周期、上下文的管理，负责爬虫的运行。
* `schedule`:负责爬虫请求的调度。(目前只有一种基于channel的调度器，无法单个爬虫多worker运行，`可自行实现基于redis，或者mq服务的调度器即可实现`)<br>
* `process (page process)`：负责请求结果的处理。<br>
* `pipline`： 结果的输出输出到不同渠道,如控制台，文件，消息队列，数据库等等<br>
* `register`：负责服务的注册（目前只支持etcd)  
* `http`: 提供一些http接口

#### 2.管理部分（manage节点）:  
负责spider节点的管理，用etcd进行spider节点的发现。通过http与spider节点通讯。

## 开始使用

#### 1. Json模版
```
http接口调用
curl -d '{"id":"douban-movie","Name":"douban-movie","request":[{"url":"https://movie.douban.com/j/new_search_subjects?sort=T\u0026range=0,10\u0026tags=\u0026start={0-100,20}","method":"get","type":"","data":null,"header":null,"cookies":{"url":"","data":""},"process_name":"movie"}],"process":[{"name":"movie","reg_url":null,"type":"json","template_rule":{"Rule":null},"json_rule":{"Rule":{"casts":"casts","cover":"cover","id":"id","node":"array|data","rate":"rate","star":"star","title":"title","url":"url"}},"add_queue":null}],"pipline":"file","depth":0,"end_count":0}' "http://127.0.0.1:7774/task/addAndRun"
```

豆瓣电影模版
```
 {
    "id": "douban-movie",
    "Name": "douban-movie",
    "request": [
        {
            "url": "https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-10,20}",
            "method": "get",
            "process_name": "movie"
        }
    ],
    "process": [
        {
            "name": "movie",
            "type": "json",
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

```
dilidili模版
``` 
   {
    "id": "dilidili",
    "Name": "dilidili",
    "request": [
        {
            "url": "http://www.dilidili.wang/{gaoxiao|kehuan|yundong|danmei|zhiyuxi|luoli|zhenren|zhuangbi|youxi|tuili|qingchun|kongbu|jizhan|rexue|qingxiaoshuo|maoxian|hougong|qihuan|tongnian|lianai|meishaonv|lizhi|baihe|paomianfan|yinv}/",
            "method": "get",
            "process_name": "animelist"
        }
    ],
    "process": [
        {
            "name": "animelist",
            "type": "template",
            "template_rule": {
                "Rule": {
                    "content": "text|dd div",
                    "desc": "text|dd p",
                    "href": "attr.href|dt a",
                    "img": "attr.src|dt a img",
                    "node": "array|.anime_list dl",
                    "title": "text|dd h3 a"
                }
            },
            "add_queue": [
                {
                    "url": "http://www.dilidili.wang{href}",
                    "method": "get",
                    "process_name": "animeinfo"
                }
            ]
        },
        {
            "name": "animeinfo",
            "type": "template",
            "template_rule": {
                "Rule": {
                    "episode": "texts|.time_con .swiper-slide .clear li a em",
                    "episode-link": "attrs.href|.time_con .swiper-slide .clear li a",
                    "title": "text|.detail dl dd h1"
                }
            },
            "add_queue": [
                {
                    "url": "{episode-link}",
                    "method": "get",
                    "process_name": "episodeinfo"
                }
            ]
        },
        {
            "name": "episodeinfo",
            "reg_url": null,
            "type": "template",
            "template_rule": {
                "Rule": {
                    "player": "attr.src|.player_main iframe",
                    "title": "text|#intro2 h1",
                    "url": "attr.href|link[rel=\"canonical\"]"
                }
            },
            "add_queue": null
        }
    ],
    "pipline": "file",
    "depth": 0,
    "end_count": 0
}
```

#### 2. 代码模版 编写
豆瓣电影
```
package main

import (
	"YiSpider/spider/model"
	"YiSpider/spider"
	spider2 "YiSpider/spider/spider"
)

func main(){

	task := &model.Task{
		Id:"douban-movie",
		Name:"douban-movie",
		Request:[]*model.Request{
			{
				Method:"get",
				Url:"https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-10000,20}",
				ProcessName:"movie",
			},
		},
		Process: []model.Process{
			{
				Name:"movie",
				Type:"json",
				JsonRule:model.JsonRule{
					Rule:map[string]string{
						"node":"array|data",
						"rate":"rate",
						"star":"star",
						"id":"id",
						"url":"url",
						"title":"title",
						"cover":"cover",
						"casts":"casts",
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
```
dilidili番剧
```
package main

import (
	"YiSpider/spider/model"
	"YiSpider/spider"
	spider2 "YiSpider/spider/spider"
)

func main(){

	task := &model.Task{
		Id:"dilidili",
		Name:"dilidili",
		Request:[]*model.Request{
			{
				Method:"get",
				Url:"http://www.dilidili.wang/{gaoxiao|kehuan|yundong|danmei|zhiyuxi|luoli|zhenren|zhuangbi|youxi|tuili|qingchun|kongbu|jizhan|rexue|qingxiaoshuo|maoxian|hougong|qihuan|tongnian|lianai|meishaonv|lizhi|baihe|paomianfan|yinv}/",
				ProcessName:"animelist",
			},
		},
		Process: []model.Process{
			{
				Name:"animelist",
				Type:"template",
				TemplateRule:model.TemplateRule{
					Rule:map[string]string{
						"node":"array|.anime_list dl",
						"img":"attr.src|dt a img",
						"title":"text|dd h3 a",
						"href":"attr.href|dt a",
						"content":"text|dd div",
						"desc":"text|dd p",
					},
				},
				AddQueue:[]*model.Request{
					{
						Method:      "get",
						Url:         "http://www.dilidili.wang{href}",
						ProcessName: "animeinfo",
					},
				},
			},
			{
				Name:"animeinfo",
				Type:"template",
				TemplateRule:model.TemplateRule{
					Rule:map[string]string{
						"episode":"texts|.time_con .swiper-slide .clear li a em",
						"title":"text|.detail dl dd h1",
						"episode-link":"attrs.href|.time_con .swiper-slide .clear li a",
					},
				},
				AddQueue:[]*model.Request{
					{
						Method:      "get",
						Url:         "{episode-link}",
						ProcessName: "episodeinfo",
					},
				},
			},
			{
				Name:"episodeinfo",
				Type:"template",
				TemplateRule:model.TemplateRule{
					Rule:map[string]string{
						"url":"attr.href|link[rel=\"canonical\"]",
						"title":"text|#intro2 h1",
						"player":"attr.src|.player_main iframe",
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
```

3. 纯代码编写
```

```
