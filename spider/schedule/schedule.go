package schedule

import (
	"YiSpider/spider/config"
	"YiSpider/spider/model"
)

type Schedule interface {
	Push(req *model.Request)
	PushMuti(reqs []*model.Request)
	Pop() (*model.Request, bool)
	Count() int
	Close()
}

var (
	scheduleMap = make(map[string]func(*config.Config) Schedule)
)

func RegisterSchedule(name string, builder func(*config.Config) Schedule) {
	scheduleMap[name] = builder
}

func GetSchedule(c *config.Config) Schedule {
	schedule := scheduleMap[c.ScheduleMode]
	if schedule == nil {
		return scheduleMap["chan"](c)
	}
	return schedule(c)
}
