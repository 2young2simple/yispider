package common

import (
	"YiSpider/spider/model"
	"fmt"
	"testing"
)

func TestPraseOffset(t *testing.T) {
	reqs := []*model.Request{
		{
			Method:      "get",
			Url:         "https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-10,1}&page={0-10,1}&last={1|2|3}",
			ProcessName: "movie",
		},
	}
	results := PraseReq(reqs, nil)
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestPraseOr(t *testing.T) {
	reqs := []*model.Request{
		{
			Method:      "get",
			Url:         "https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0|20|40}&page={1|2|3}",
			ProcessName: "movie",
		},
	}
	results := PraseReq(reqs, nil)
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestPraseParamCtx(t *testing.T) {
	reqs := []*model.Request{
		{
			Method:      "get",
			Url:         "https://sclub.jd.com/comment/productPageComments.action?productId={$id}&score=0&sortType=5&page={0-$max_page,1}&pageSize=10",
			ProcessName: "movie",
		},
	}
	results := PraseReq(reqs, map[string]interface{}{
		"id":       13123123,
		"max_page": 10,
	})
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestFindRule(t *testing.T) {
	url := `"https://movie.douban.com/j/new_search_subjects?sort=T&url={1-$count,1}&tags="`

	results := FindRule(url)
	for _, result := range results {
		fmt.Println(result)
	}
}
