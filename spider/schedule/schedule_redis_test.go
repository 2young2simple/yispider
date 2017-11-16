package schedule

import (
	"testing"

	"YiSpider/spider/config"
	"YiSpider/spider/model"
)

func TestRedisSchedule_Push(t *testing.T) {
	s := NewRedisSchedule(&config.Config{RedisAddr: "127.0.0.1:6379"})
	s.Push(&model.Request{Url: "www.bai123.com", Method: "get", Header: map[string]string{"a": "b"}})
}

func TestRedisSchedule_Pop(t *testing.T) {
	s := NewRedisSchedule(&config.Config{Name: "qiongyou_spider", RedisAddr: "127.0.0.1:6379"})
	for i := 0; i < 100; i++ {
		go s.Pop()
	}
}
