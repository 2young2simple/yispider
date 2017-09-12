package filter

import (
	"testing"
	"YiSpider/spider/model"
	"fmt"
)

func TestFilter(t *testing.T) {
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

	fmt.Println(Filter("/8hr/",task))
	fmt.Println(Filter("/8hr/page/3/",task))
	fmt.Println(Filter("/8hr/page/4/",task))
	fmt.Println(Filter("/8hr/page/5/",task))
	fmt.Println(Filter("/8hr/page/13/",task))
	fmt.Println(Filter("/8hr/page/3/",task))
}
