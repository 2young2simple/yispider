package downloader

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"net/http/cookiejar"
	"golang.org/x/net/publicsuffix"
	"time"
	"encoding/json"
	"YiSpider/spider/logger"
)

var Clients map[string]*http.Client

func init(){
	Clients = make(map[string]*http.Client)
}

func makeCookiejar() http.CookieJar {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&cookiejarOptions)

	return jar
}

func makeClient(transport http.RoundTripper, jar http.CookieJar) *http.Client {
	return &http.Client{Jar: jar, Transport: transport, Timeout: 60 * time.Second}
}

func Get(taskId string,url string) ([]byte,error) {
	res, err := doRequest(taskId,"GET",url,nil)
	if err != nil {
		return nil,err
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return nil,err
	}
	defer res.Body.Close()
	logger.Info("GET",url, " =>", string(body))
	return body,nil
}

func PostJson(taskId string,url string,data interface{}) ([]byte,error) {
	dataJ,err := json.Marshal(data)
	if err != nil{
		return nil,err
	}
	logger.Info("Request:",string(dataJ))
	res, err := doRequest(taskId,"POST",url,dataJ)
	if err != nil {
		return nil,err
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return nil,err
	}
	logger.Info("POST",url,"=>",string(body))
	return body,nil
}

func doRequest(id string,method string,url string,data []byte) (resp *http.Response, err error){
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := Clients[id]
	if client == nil{
		client = makeClient(nil,makeCookiejar())
		Clients[id] = client
	}
	return client.Do(req)
}
