package http

import (
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"YiSpider/manage/model"
)

var errorMethod = []byte("{\"code\":\"400\",\"msg\":\"not support method\"}")
var errorQuery = []byte("{\"code\":\"400\",\"msg\":\"error url parmas\"}")
var errorBody = []byte("{\"code\":\"400\",\"msg\":\"error get body\"}")
var errorJson = []byte("{\"code\":\"400\",\"msg\":\"error get Json\"}")
var commonSuccess = []byte("{\"code\":\"200\",\"msg\":\"success\"}")



func AddTask(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST"{
		w.Write(errorMethod)
		return
	}
	body,err := ioutil.ReadAll(req.Body)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	task := &model.Task{}
	err = json.Unmarshal(body,task)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	data,err := AddTaskS(task)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
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

	data,err := StopTaskS(name)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
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

	data,err := RunTaskS(name)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
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

	data,err := EndTaskS(name)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}

func ListTask(w http.ResponseWriter, req *http.Request) {
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

	data,err := ListTaskS(name)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}

func ListNode(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET"{
		w.Write(errorMethod)
		return
	}

	data,err := ListNodesS()
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(data)
}
