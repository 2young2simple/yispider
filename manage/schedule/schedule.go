package schedule

import (
	"YiSpider/manage/discover"
	"YiSpider/manage/model"
	"YiSpider/manage/strategy"
	"fmt"
)

func AddTask(task *model.Task) ([]byte, error) {
	node := strategy.GetNode()
	return Post(getUrl(node.IP, "/task/add"), task)
}

func RunTask(name string) ([]byte, error) {
	node := strategy.GetNode()
	return Get(getUrl(node.IP, "/task/run?name="+name))
}

func StopTask(name string) ([]byte, error) {
	node := strategy.GetNode()
	return Get(getUrl(node.IP, "/task/stop?name="+name))
}

func EndTask(name string) ([]byte, error) {
	node := strategy.GetNode()
	return Get(getUrl(node.IP, "/task/end?name="+name))
}

func ListTask(name string) ([]byte, error) {
	node := discover.GetNodes()[name]
	return Get(getUrl(node.IP, "/task/list"))
}

func getUrl(ip string, path string) string {
	url := fmt.Sprintf("http://%s:7777%s", ip, path)
	return url
}
