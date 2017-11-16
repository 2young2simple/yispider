package downloader

import (
	"YiSpider/spider/logger"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"
)

var Clients map[string]*http.Client
var lock sync.RWMutex

func init() {
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

func Get(taskId string, url string) (*http.Response, error) {
	res, err := doRequest(taskId, "GET", url, nil)
	if err != nil {
		logger.Info("Download fail doRequest,url:", url, "err:", err)
		return nil, err
	}
	logger.Info("GET", url, " =>", res.StatusCode)
	if res.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("download fail,url %s, StatusCode %d", url, res.StatusCode))
	}
	return res, nil
}

func PostJson(taskId string, url string, data interface{}) (*http.Response, error) {
	dataJ, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := doRequest(taskId, "POST", url, dataJ)
	if err != nil {
		return nil, err
	}
	logger.Info("POST", url, "=>", res.StatusCode)
	if res.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("download fail, StatusCode %d", res.StatusCode))
	}
	return res, nil
}

func doRequest(id string, method string, url string, data []byte) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type", "application/json")
	client := getClient(id)
	if client == nil {
		client = makeClient(nil, makeCookiejar())
		setClient(id, client)
	}
	return client.Do(req)
}

func setClient(id string, client *http.Client) {
	lock.Lock()
	defer lock.Unlock()
	Clients[id] = client
}

func getClient(id string) *http.Client {
	lock.RLock()
	defer lock.RUnlock()
	client := Clients[id]
	return client
}
