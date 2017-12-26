package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
)

func main() {

	task := &model.Task{
		Id:   "ttkb-author",
		Name: "ttkb-author",
		Request: []*model.Request{
			{
				Method:      "get",
				Url:`http://r.cnews.qq.com/getSubItem?chlid={6000000-6200000,1}`,
				ProcessName: "ttkb-author",
			},
		},
		Process: []model.Process{
			{
				Name: "ttkb-author",
				Type: "json",
				JsonRule: model.JsonRule{
					Rule: map[string]string{
						"chlid":    "channelInfo.chlid",
						"chlname":"channelInfo.chlname",
						"desc":"channelInfo.desc",
						"subCount":"channelInfo.subCount",
						"uin":"channelInfo.uin",
						"intro":"channelInfo.intro",
						"recommend":"channelInfo.recommend",
						"followCount":"channelInfo.followCount",
						"readCount":"channelInfo.readCount",
						"shareCount":"channelInfo.shareCount",
						"colCount":"channelInfo.colCount",
					},
				},
			},
		},

		Pipline: "mysql",
	}

	app := spider.New()
	app.AddSpider(spider2.InitWithTask(task))
	app.Run()
}
