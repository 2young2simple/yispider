package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
	"fmt"
)

func main() {
	types := "daily_timeline|kb_video_news|kb_news_bagua|kb_news_qipa|kb_photo_news|kb_news_tech|kb_news_finance|location|kb_news_world|kb_news_movie|kb_news_gaojidi|kb_news_wealth|kb_photo_gif|kb_news_sports|kb_news_mil|kb_news_history|kb_news_nba|kb_news_car|kb_news_chaobao|kb_news_laugh|kb_news_pet|kb_news_science|kb_news_baby|kb_news_astro|kb_news_sex|kb_news_beauty|kb_news_house|kb_news_share|kb_news_rock|kb_news_tfboys|kb_news_augury|kb_news_photography|kb_news_lottery|kb_news_cate|kb_news_julebu|kb_news_travel|kb_news_idea|kb_news_lol|kb_news_erciyuan|kb_news_space|kb_news_game|kb_news_iphone|kb_news_esport|kb_news_health|kb_news_outfit|kb_news_furnishing|kb_news_workout|kb_news_soup|kb_news_run|kb_news_fishing|kb_news_buddism|kb_news_diet|kb_news_football|kb_news_tennis|kb_news_tea|kb_news_yoga|kb_news_plaything|kb_news_watch"
	//types := "daily_timeline|kb_video_news|kb_news_bagua|kb_news_qipa|kb_photo_news|kb_news_tech|kb_news_finance|location|kb_news_world|kb_news_movie|kb_news_gaojidi|kb_news_wealth|kb_photo_gif|kb_news_sports|kb_news_mil|kb_news_history|kb_news_nba|kb_news_car|kb_news_chaobao|kb_news_laugh|kb_news_pet|kb_news_science|kb_news_baby|kb_news_astro|kb_news_sex|kb_news_beauty|kb_news_house|kb_news_share|kb_news_rock|kb_news_tfboys|kb_news_augury|kb_news_photography|kb_news_lottery|kb_news_cate|kb_news_julebu|kb_news_travel|kb_news_idea|kb_news_lol|kb_news_erciyuan|kb_news_space|kb_news_game|kb_news_iphone|kb_news_esport|kb_news_health|kb_news_outfit|kb_news_furnishing|kb_news_workout|kb_news_soup|kb_news_run|kb_news_fishing|kb_news_buddism|kb_news_diet|kb_news_football|kb_news_tennis|kb_news_tea|kb_news_yoga|kb_news_plaything|kb_news_watch"

	task := &model.Task{
		Id:   "ttkb-author",
		Name: "ttkb-author",
		Request: []*model.Request{
			{
				Method:      "get",
				Url:         fmt.Sprintf(`http://r.cnews.qq.com/getSubNewsChlidInterest?patchver=4511&mid=fd248c13ee1ce793495484e4cf3250f8ebbb475a&devid=860046037899335&store=60009&screen_height=1920&apptype=android&origin_imei=860046037899335&hw=OnePlus_ONEPLUSA3000&appver=25_areading_4.5.11&appversion=4.5.11&uid=bfa0a264a6547298&screen_width=1080&sceneid=&android_id=bfa0a264a6547298&last_id=20171207A03G7J00&ssid=GeeyueTech_5G&forward=0&IronThroneBuildTime=1512716487405&omgid=e0f7a4180378ba4e5ee80b0820ef5a1744ca0010211815&IronThroneRelBuildTime=415047497&refreshType=normal&qqnetwork=wifi&last_time=&bottom_id=20171207A0BFU500&top_time=1512631500&currentTab=kuaibao&top_id=20171207C0HX4500&is_wap=0&omgbizid=b03081d3f5806f45b65904d08cfad6bc77130080211815&page={1-1000,1}&imsi=460019017167485&lastRefreshTime=&IronThroneRelExecTime=415047499&muid=49887860909485482&activefrom=icon&cachedCount=20&direction=0&sessionid=&chRefreshTimes=0&chlid={%s}&bottom_time=1512603257&IronThroneExecTime=1512716487407&qn-sig=284d6905ece4010e0ebd89dce072b5ee&qn-rid=6e63ca4d-1285-47ee-b95d-0bb49da3ce03`,types),
				ProcessName: "ttkblist",
			},
		},
		Process: []model.Process{
			{
				Name: "ttkblist",
				Type: "json",
				JsonRule: model.JsonRule{
					Rule: map[string]string{
						"node":   "array|newslist",
						"chlid":    "chlid",
					},
				},
				AddQueue:[]*model.Request{
					{
						Method:  "get",
						Url :    "http://r.cnews.qq.com/getSubItem?chlid={$chlid}",
						ProcessName: "author",
					},
				},
			},
			{
				Name: "author",
				Type: "json",
				JsonRule: model.JsonRule{
					Rule: map[string]string{
						//"node":   "",
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
