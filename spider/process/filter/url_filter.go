package filter

import (
	"YiSpider/spider/model"
	"regexp"
)

func Filter(url string, process *model.Process) bool {
	if len(url) == 0 {
		return false
	}

	check := false
	for _, regUrl := range process.RegUrl {
		reg := regexp.MustCompile(regUrl)
		match := reg.MatchString(url)
		if match {
			check = true
			break
		}
	}

	if check == false {
		return false
	}

	return RepeatFilter(url, process)
}
