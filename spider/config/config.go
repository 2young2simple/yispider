package config

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"YiSpider/spider/logger"
)

var ConfigI *Config

type Config struct {
	Name string `json:"name"`
	Version string `json:"version"`
	WorkNum int `json:"work_num"`
	MaxWaitNum int `json:"max_wait_num"`

	HttpAddr string `json:"http_addr"`
	RedisAddr string `json:"redis_addr"`
	ScheduleMode string `json:"schedule"`
	Etcd []string `json:"etcd"`
}

func InitConfig() error{

	var file *os.File
	var bytes []byte
	var err error

	if file,err = os.OpenFile("conf.json",os.O_RDONLY,0666);err != nil{
		return err
	}

	if bytes,err = ioutil.ReadAll(file);err != nil{
		return err
	}

	ConfigI = &Config{}
	if err = json.Unmarshal(bytes,ConfigI);err != nil{
		return err
	}

	logger.Info("init success ",*ConfigI)
	return nil
}