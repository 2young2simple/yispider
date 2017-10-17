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


/*
   dilidili json

   {
    "id": "dilidili",
    "Name": "dilidili",
    "request": [
        {
            "url": "http://www.dilidili.wang/{gaoxiao|kehuan|yundong|danmei|zhiyuxi|luoli|zhenren|zhuangbi|youxi|tuili|qingchun|kongbu|jizhan|rexue|qingxiaoshuo|maoxian|hougong|qihuan|tongnian|lianai|meishaonv|lizhi|baihe|paomianfan|yinv}/",
            "method": "get",
            "type": "",
            "data": null,
            "header": null,
            "cookies": {
                "url": "",
                "data": ""
            },
            "process_name": "animelist"
        }
    ],
    "process": [
        {
            "name": "animelist",
            "reg_url": null,
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
            "json_rule": {
                "Rule": null
            },
            "add_queue": [
                {
                    "url": "http://www.dilidili.wang{href}",
                    "method": "get",
                    "type": "",
                    "data": null,
                    "header": null,
                    "cookies": {
                        "url": "",
                        "data": ""
                    },
                    "process_name": "animeinfo"
                }
            ]
        },
        {
            "name": "animeinfo",
            "reg_url": null,
            "type": "template",
            "template_rule": {
                "Rule": {
                    "episode": "texts|.time_con .swiper-slide .clear li a em",
                    "episode-link": "attrs.href|.time_con .swiper-slide .clear li a",
                    "title": "text|.detail dl dd h1"
                }
            },
            "json_rule": {
                "Rule": null
            },
            "add_queue": [
                {
                    "url": "{episode-link}",
                    "method": "get",
                    "type": "",
                    "data": null,
                    "header": null,
                    "cookies": {
                        "url": "",
                        "data": ""
                    },
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
            "json_rule": {
                "Rule": null
            },
            "add_queue": null
        }
    ],
    "pipline": "file",
    "depth": 0,
    "end_count": 0
}

{"id":"dilidili","Name":"dilidili","request":[{"url":"http://www.dilidili.wang/{gaoxiao|kehuan|yundong|danmei|zhiyuxi|luoli|zhenren|zhuangbi|youxi|tuili|qingchun|kongbu|jizhan|rexue|qingxiaoshuo|maoxian|hougong|qihuan|tongnian|lianai|meishaonv|lizhi|baihe|paomianfan|yinv}/","method":"get","type":"","data":null,"header":null,"cookies":{"url":"","data":""},"process_name":"animelist"}],"process":[{"name":"animelist","reg_url":null,"type":"template","template_rule":{"Rule":{"content":"text|dd div","desc":"text|dd p","href":"attr.href|dt a","img":"attr.src|dt a img","node":"array|.anime_list dl","title":"text|dd h3 a"}},"json_rule":{"Rule":null},"add_queue":[{"url":"http://www.dilidili.wang{href}","method":"get","type":"","data":null,"header":null,"cookies":{"url":"","data":""},"process_name":"animeinfo"}]},{"name":"animeinfo","reg_url":null,"type":"template","template_rule":{"Rule":{"episode":"texts|.time_con .swiper-slide .clear li a em","episode-link":"attrs.href|.time_con .swiper-slide .clear li a","title":"text|.detail dl dd h1"}},"json_rule":{"Rule":null},"add_queue":[{"url":"{episode-link}","method":"get","type":"","data":null,"header":null,"cookies":{"url":"","data":""},"process_name":"episodeinfo"}]},{"name":"episodeinfo","reg_url":null,"type":"template","template_rule":{"Rule":{"player":"attr.src|.player_main iframe","title":"text|#intro2 h1","url":"attr.href|link[rel=\"canonical\"]"}},"json_rule":{"Rule":null},"add_queue":null}],"pipline":"file","depth":0,"end_count":0}

 */