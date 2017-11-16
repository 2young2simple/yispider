package discover

import (
	"YiSpider/manage/config"
	"YiSpider/manage/discover/etcd"
	"YiSpider/manage/model"
)

type Discover interface {
	GetNodes() map[string]*model.Node
	Start() error
}

var DiscoverI Discover

func InitDiscover() error {
	var err error
	switch config.ConfigI.Discover {
	case "etcd":
		DiscoverI, err = etcd.NewCluster(config.ConfigI.Etcd)
		if err != nil {
			return err
		}
		DiscoverI.Start()
	}
	return nil
}

func GetNodes() map[string]*model.Node {
	if DiscoverI != nil {
		return DiscoverI.GetNodes()
	}
	return nil
}
