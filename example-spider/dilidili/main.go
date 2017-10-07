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
						ProcessName: "animeinfo",
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
