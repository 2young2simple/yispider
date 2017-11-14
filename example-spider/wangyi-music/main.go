package main

import (
	"YiSpider/spider/model"
	"YiSpider/spider"
	spider2 "YiSpider/spider/spider"
	"strings"
	"fmt"
)

func main(){
	musicType := `华语| 欧美| 日语| 韩语| 粤语| 小语种|`
////流行| 摇滚| 民谣| 电子| 舞曲| 说唱| 轻音乐| 爵士| 乡村| R&B/Soul| 古典| 民族| 英伦| 金属| 朋克| 蓝调| 雷鬼| 世界音乐| 拉丁| 另类/独立| New Age| 古风| 后摇| Bossa Nova|
////清晨| 夜晚| 学习| 工作| 午休| 下午茶| 地铁| 驾车| 运动| 旅行| 散步| 酒吧|
////怀旧| 清新| 浪漫| 性感| 伤感| 治愈| 放松| 孤独| 感动| 兴奋| 快乐| 安静| 思念|
////影视原声| ACG| 校园| 游戏| 70后| 80后| 90后| 网络歌曲| KTV| 经典| 翻唱| 吉他| 钢琴| 器乐| 儿童| 榜单| 00后|`
//	musicType := `华语`
	musicType = strings.Replace(musicType," ","",-1)
	musicTypes := strings.Split(musicType,"|")
	reqs := []*model.Request{}

	for _,ty := range musicTypes{
		reqs = append (reqs,&model.Request{
			Method:"get",
			Url:fmt.Sprintf("http://music.163.com/discover/playlist/?order=hot&cat=%s&limit=35&offset={0-1440,35}",ty),
			ProcessName:"music-list",
		})
	}
	task := &model.Task{
		Id:"music-list",
		Name:"music-list",
		Request:reqs,
		Process: []model.Process{
			{
				Name:"music-list",
				Type:"template",
				TemplateRule:model.TemplateRule{
					Rule:map[string]string{
						"node":"array|.m-cvrlst li",
						"img":"attr.src|.u-cover img",
						"music_addr":"attr.href|.u-cover a",
						"title":"attr.title|.u-cover a",
						"play_num":"text|.nb",
						"author":"text|.nm",
					},
				},
				AddQueue:[]*model.Request{
					{
						Method:"get",
						Url:"http://music.163.com{music_addr}",
						ProcessName:"music-detail",
					},
				},
			},
			{
				Name:"music-detail",
				Type:"template",
				TemplateRule:model.TemplateRule{
					Rule:map[string]string{
						"img":"attr.src|.u-cover img",
						"title":"text|.f-ff2",
						"play_num":"text|#play-count",
						"author":"text|.s-fc7",
						"like_num":"text|.u-btni-fav i",
						"share_num":"text|.u-btni-share i",
						"comment_num":"text|#cnt_comment_count",
						"desc":"text|#album-desc-dot",
						"time":"text|.time",
						"music_count":"#playlist-track-count",
						"id":"attr.data-rid|#content-operation",
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
