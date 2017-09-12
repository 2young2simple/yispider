package filter

import (
	"YiSpider/spider/model"
	"regexp"
	"strings"
)

func Filter(url string,task *model.Task) bool{
	if len(url) == 0 {
		return false
	}
	if !strings.HasPrefix(url,"/"){
		return false
	}

	check := false
	for _,regUrl := range task.Process.RegUrl{
		reg := regexp.MustCompile(regUrl)
		match := reg.MatchString(url)
		if match{
			check = true
			break
		}
	}

	if check == false{
		return false
	}

	return RepeatFilter(url,task)
}
