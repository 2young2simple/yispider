package filter

import (
	"YiSpider/spider/model"
)

var CuckooFilter map[string]int

func init(){
	CuckooFilter = make(map[string]int)
}

func RepeatFilter(url string,task *model.Task) bool{
	sign := task.Name+":"+url
	 if _,ok := CuckooFilter[sign];ok{
		return false
	}

	CuckooFilter[sign] = 1

	return true
}

