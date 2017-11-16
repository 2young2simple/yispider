package http

import (
	"YiSpider/manage/config"
	"YiSpider/manage/logger"
	"net/http"
)

func InitHttpServer() {

	http.HandleFunc("/task/add", AddTask)
	http.HandleFunc("/task/run", RunTask)
	http.HandleFunc("/task/stop", StopTask)
	http.HandleFunc("/task/end", EndTask)
	http.HandleFunc("/tasks", ListTask)
	http.HandleFunc("/nodes", ListNode)

	err := http.ListenAndServe(config.ConfigI.HttpAddr, nil)
	if err != nil {
		logger.Error("ListenAndServe fail:", err)
	}
}
