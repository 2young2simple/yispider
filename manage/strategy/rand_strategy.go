package strategy

import (
	"YiSpider/manage/model"
	"YiSpider/manage/discover"
)

func GetNode() *model.Node{
	nodes := discover.GetNodes()
	for _,node := range nodes{
		return node
	}
	return nil
}