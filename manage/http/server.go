package http

import (
	"net/http"
	"io"
	"log"
	"net/url"
	"YiSpider/spider/core"
	"encoding/json"
)

var errorMethod = []byte("{\"code\":\"400\",\"msg\":\"not support method\"}")
var errorQuery = []byte("{\"code\":\"400\",\"msg\":\"error url parmas\"}")
var commonSuccess = []byte("{\"code\":\"200\",\"msg\":\"success\"}")



func AddTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST"{
		w.Write(errorMethod)
		return
	}
	io.WriteString(w, "hello, world!\n")
}

func StopTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET"{
		w.Write(errorMethod)
		return
	}

	queryMap, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil{
		w.Write(errorQuery)
		return
	}
	name := queryMap.Get("name")
	if err := core.GetEnine().StopTask(name);err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(commonSuccess)
	return
}

func RunTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET"{
		w.Write(errorMethod)
		return
	}

	queryMap, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil{
		w.Write(errorQuery)
		return
	}
	name := queryMap.Get("name")
	if err := core.GetEnine().RunTask(name);err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(commonSuccess)
	return
}

func EndTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET"{
		w.Write(errorMethod)
		return
	}

	queryMap, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil{
		w.Write(errorQuery)
		return
	}
	name := queryMap.Get("name")
	if err := core.GetEnine().EndTask(name);err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(commonSuccess)
	return
}

func ListTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET"{
		w.Write(errorMethod)
		return
	}

	tasks := core.GetEnine().ListTask()
	datas,err := json.Marshal(tasks)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(datas)
	return
}

func InitHttpServer() {
	http.HandleFunc("/task/add", AddTask)
	http.HandleFunc("/task/run", RunTask)
	http.HandleFunc("/task/stop", StopTask)
	http.HandleFunc("/task/end", EndTask)
	http.HandleFunc("/tasks", ListTask)


	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("ListenAndServe fail:", err)
	}
}