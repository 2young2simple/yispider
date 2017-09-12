package core

import (
	"github.com/kataras/go-errors"
	"fmt"
	"YiSpider/spider/config"
	"YiSpider/spider/logger"
	"YiSpider/spider/model"
)

var engineI *Engine

func init(){
	var err error
	if err = config.InitConfig();err != nil{
		logger.Info(err.Error())
	}
}

func New() *Engine{
	engineI = &Engine{spiders:make(map[string]*Spider)}
	return engineI
}

func GetEnine() *Engine{
	return engineI
}


type Engine struct {
	spiders map[string]*Spider
}

func (m *Engine)AddTask(task *model.Task){
	spider := NewSpiderTask()
	spider.SetTask(task)
	m.spiders[spider.task.Name] = spider
}

func (m *Engine)RunTask(name string) error{
	spider,ok := m.spiders[name]
	if !ok{
		return errors.New(fmt.Sprintf("Task [%s] is not exist",name))
	}
	spider.Run()
	return nil
}

func (m *Engine)StopTask(name string) error{
	spider,ok := m.spiders[name]
	if !ok{
		return errors.New(fmt.Sprintf("Task [%s] is not exist",name))
	}
	spider.Stop()
	return nil
}

func (m *Engine)EndTask(name string) error{
	spider,ok := m.spiders[name]
	if !ok{
		return errors.New(fmt.Sprintf("Task [%s] is not exist",name))
	}
	spider.Exit()
	return nil
}

func (m *Engine) ListTask() []*model.Task{
	spiders := []*model.Task{}
	for _,spider := range m.spiders{
		spiders = append(spiders,spider.GetTask())
	}
	return spiders
}


func (m *Engine) Run() {
	for _,spider := range m.spiders{
		spider.Run()
	}
}



