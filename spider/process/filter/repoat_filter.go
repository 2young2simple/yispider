package filter

import (
	"YiSpider/spider/model"
	"sync"
)

var CuckooFilter map[string]int
var lock sync.RWMutex

func init() {
	CuckooFilter = make(map[string]int)
}

func RepeatFilter(url string, process *model.Process) bool {
	sign := url
	if ok := get(sign); ok {
		return false
	}
	put(sign)
	return true
}

func get(name string) bool {
	lock.RLock()
	defer lock.RUnlock()
	_, ok := CuckooFilter[name]
	return ok
}

func put(name string) {
	lock.Lock()
	defer lock.Unlock()
	CuckooFilter[name] = 1
}
