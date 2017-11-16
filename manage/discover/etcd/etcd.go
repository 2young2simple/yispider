package etcd

import (
	"encoding/json"
	"time"

	"YiSpider/manage/logger"
	"YiSpider/manage/model"
	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"log"
)

type Cluster struct {
	nodes   map[string]*model.Node
	KeysAPI client.KeysAPI
}

func NewCluster(endpoints []string) (*Cluster, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		logger.Error("Error: cannot connec to etcd:", err)
		return nil, err
	}

	master := &Cluster{
		nodes:   make(map[string]*model.Node),
		KeysAPI: client.NewKeysAPI(etcdClient),
	}
	return master, nil
}

func (c *Cluster) Start() error {
	go c.WatchWorkers()
	fmt.Println("Master Start ...")
	return nil
}

func (c *Cluster) GetNodes() map[string]*model.Node {
	fmt.Println("c.nodes", c.nodes)
	return c.nodes
}

func (c *Cluster) addWorker(info *model.WorkerInfo) {
	node := &model.Node{
		IsHealth:   true,
		IP:         info.IP,
		Name:       info.Name,
		CPU:        info.CPU,
		MetaData:   info.MetaData,
		SpiderData: info.SpiderData,
	}
	c.nodes[node.Name] = node
}

func (c *Cluster) updateWorker(info *model.WorkerInfo) {
	c.addWorker(info)
}

func unmarshal(node *client.Node) *model.WorkerInfo {
	logger.Info(node.Value)
	info := &model.WorkerInfo{}
	err := json.Unmarshal([]byte(node.Value), info)
	if err != nil {
		logger.Error(err)
	}
	return info
}

func (c *Cluster) WatchWorkers() {
	api := c.KeysAPI
	watcher := api.Watcher("spiders/", &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			logger.Error("Error watch workers:", err)
			break
		}
		if res.Action == "expire" {
			info := unmarshal(res.PrevNode)
			logger.Info("Expire worker ", info.Name)
			member, ok := c.nodes[info.Name]
			if ok {
				member.IsHealth = false
			}
		} else if res.Action == "set" {
			info := unmarshal(res.Node)
			if _, ok := c.nodes[info.Name]; ok {
				logger.Info("Update worker ", info.Name)
				c.updateWorker(info)
			} else {
				logger.Info("Add worker ", info.Name)
				c.addWorker(info)
			}
		} else if res.Action == "delete" {
			info := unmarshal(res.Node)
			log.Println("Delete worker ", info.Name)
			delete(c.nodes, info.Name)
		}
	}
}
