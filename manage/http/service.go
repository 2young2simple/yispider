package http

import (
	"YiSpider/manage/discover"
	"YiSpider/manage/model"
	"YiSpider/manage/strategy"
	"encoding/json"
	"fmt"
)

func AddTaskS(task *model.Task) ([]byte, error) {
	node := strategy.GetNode()
	return Post(getUrl(node.IP, "/task/add"), task)
}

func RunTaskS(name string) ([]byte, error) {
	node := strategy.GetNode()
	return Get(getUrl(node.IP, "/task/run?name="+name))
}

func StopTaskS(name string) ([]byte, error) {
	node := strategy.GetNode()
	return Get(getUrl(node.IP, "/task/stop?name="+name))
}

func EndTaskS(name string) ([]byte, error) {
	node := strategy.GetNode()
	return Get(getUrl(node.IP, "/task/end?name="+name))
}

func ListTaskS(name string) ([]byte, error) {
	fmt.Println("name", name, "nodes", discover.GetNodes())
	node := discover.GetNodes()[name]
	return Get(getUrl(node.IP, "/tasks"))
}

func ListNodesS() ([]byte, error) {
	nodes := discover.GetNodes()
	return json.Marshal(nodes)
}

func getUrl(ip string, path string) string {
	url := fmt.Sprintf("http://%s%s", ip, path)
	return url
}
