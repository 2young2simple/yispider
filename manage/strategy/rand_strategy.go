package strategy

import (
	"YiSpider/manage/discover"
	"YiSpider/manage/model"
)

func GetNode() *model.Node {
	nodes := discover.GetNodes()
	for _, node := range nodes {
		return node
	}
	return nil
}
