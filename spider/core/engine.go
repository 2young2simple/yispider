package core

import (
	"YiSpider/spider/spider"
	"fmt"
	"github.com/kataras/go-errors"
	"sync"
)

var engineI *Engine
var once sync.Once

func New() *Engine {
	once.Do(func() {
		engineI = &Engine{spiders: make(map[string]*SpiderRuntime)}
	})
	return engineI
}

func GetEnine() *Engine {
	return engineI
}

type Engine struct {
	spiders map[string]*SpiderRuntime
}

func (m *Engine) AddSpider(spider *spider.Spider) *Engine {
	spiderRuntime := NewSpiderRuntime()
	spiderRuntime.SetSpider(spider)
	m.spiders[spider.Name] = spiderRuntime
	return m
}

func (m *Engine) RunTask(name string) error {
	s, ok := m.spiders[name]
	if !ok {
		return errors.New(fmt.Sprintf("Task [%s] is not exist", name))
	}
	s.Run()
	return nil
}

func (m *Engine) StopTask(name string) error {
	s, ok := m.spiders[name]
	if !ok {
		return errors.New(fmt.Sprintf("Task [%s] is not exist", name))
	}
	s.Stop()
	return nil
}

func (m *Engine) EndTask(name string) error {
	s, ok := m.spiders[name]
	if !ok {
		return errors.New(fmt.Sprintf("Task [%s] is not exist", name))
	}
	s.Exit()
	return nil
}

func (m *Engine) ListTask() []*spider.Spider {
	spiders := []*spider.Spider{}
	for _, s := range m.spiders {
		spiders = append(spiders, s.spider)
	}
	return spiders
}

func (m *Engine) GetTaskMetas() map[string]*TaskMeta {
	metas := map[string]*TaskMeta{}
	for name, s := range m.spiders {
		metas[name] = s.TaskMeta
	}
	return metas
}

func (m *Engine) Run() {
	for _, s := range m.spiders {
		s.Run()
	}
}
