package schedule

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"net/http/cookiejar"
	"golang.org/x/net/publicsuffix"
	"time"
	"testing"
	"encoding/json"
)


var ClientI *http.Client

func init(){
	ClientI = MakeClient(nil)
}


func makeCookiejar() http.CookieJar {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&cookiejarOptions)

	return jar
}

func MakeClient(transport http.RoundTripper) *http.Client {
	return &http.Client{Jar: makeCookiejar(), Transport: transport, Timeout: 60 * time.Second}
}


func Get(url string) ([]byte,error) {
	res, err := DoRequest("GET",url,nil)
	if err != nil {
		return []byte{},err
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return []byte{},err
	}
	fmt.Println("GET",url, " =>", string(body))

	return body,nil
}

func Post(url string,data interface{}) ([]byte,error) {
	dataJ,err := json.Marshal(data)
	if err != nil{
		return []byte{},err
	}
	fmt.Println("Request:",string(dataJ))
	res, err := DoRequest("POST",url,dataJ)
	if err != nil {
		return []byte{},err
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return []byte{},err
	}

	return body,nil
}

func DoRequest(method string,url string,data []byte) (resp *http.Response, err error){
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return ClientI.Do(req)
}
