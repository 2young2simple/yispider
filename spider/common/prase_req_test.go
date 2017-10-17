package common

import (
	"testing"
	"YiSpider/spider/model"
	"fmt"
)

func TestPraseOffset(t *testing.T) {
	reqs := []*model.Request{
		{
			Method:"get",
			Url:"https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-400,20}",
			ProcessName:"movie",
		},
	}
	results := PraseReq(reqs,nil)
	for _,result := range results{
		fmt.Println(result)
	}
}

func TestPraseOr(t *testing.T) {
	reqs := []*model.Request{
		{
			Method:"get",
			Url:"https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0|20|40}",
			ProcessName:"movie",
		},
	}
	results := PraseReq(reqs,nil)
	for _,result := range results{
		fmt.Println(result)
	}
}

func TestPraseParamCtx(t *testing.T) {
	reqs := []*model.Request{
		{
			Method:"get",
			Url:"https://movie.douban.com/j/new_search_subjects?sort=T&url={url}&tags=&name={name}",
			ProcessName:"movie",
		},
	}
	results := PraseReq(reqs,map[string]string{
		"name":"aabb",
		"url":"uull",
	})
	for _,result := range results{
		fmt.Println(result)
	}
}