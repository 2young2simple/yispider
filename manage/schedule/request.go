package http

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

const (
	HTTP = "http://"
	HOST = "127.0.0.1:8089"
)

var TestClient *http.Client

func MakeCookiejar() http.CookieJar {
	cookiejarOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, _ := cookiejar.New(&cookiejarOptions)

	return jar
}

func MakeClient(transport http.RoundTripper, jar http.CookieJar) *http.Client {
	return &http.Client{Jar: jar, Transport: transport, Timeout: 60 * time.Second}
}

func Get(t *testing.T,url string) string {
	res, err := DoRequest("GET",url,nil)
	if err != nil {
		t.Fatal(err)
		return ""
	}

	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
		return ""
	}
	return string(body)
}

func Get_AssertCode(t *testing.T, path string, want int) string {
	url := HTTP + HOST + path
	res, err := DoRequest("GET",url,nil)
	if err != nil {
		t.Fatal(err)
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}
	fmt.Println("GET",path, " =>", string(body))
	if res.StatusCode != want {
		e := fmt.Sprintf("%s request StatusCode want %d but get %d", path, want, res.StatusCode)
		t.Fatal(e)
	}
	return string(body)
}

func Post_AssertCode(t *testing.T, path string,data interface{}, want int) string {
	dataJ,err := json.Marshal(data)
	if err != nil{
		t.Fatal("Json.Marshal err:",err)
	}
	fmt.Println("Request:",string(dataJ))
	url := HTTP + HOST + path
	res, err := DoRequest("POST",url,dataJ)
	if err != nil {
		t.Fatal(err)
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}
	fmt.Println("POST",path,"=>",string(body))
	if res.StatusCode != want {
		e := fmt.Sprintf("%s request StatusCode want %d but get %d", path, want, res.StatusCode)
		t.Fatal(e)
	}
	return string(body)
}
func PostByte_AssertCode(t *testing.T, path string,data []byte, want int) string {

	url := HTTP + HOST + path
	res, err := DoRequest("POST",url,data)
	if err != nil {
		t.Fatal(err)
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}
	fmt.Println("POST",path,"=>",string(body))
	if res.StatusCode != want {
		e := fmt.Sprintf("%s request StatusCode want %d but get %d", path, want, res.StatusCode)
		t.Fatal(e)
	}
	return string(body)
}

func Delete_AssertCode(t *testing.T, path string,data interface{}, want int) string {
	dataJ,err := json.Marshal(data)
	if err != nil{
		t.Fatal("Json.Marshal err:",err)
	}

	url := HTTP + HOST + path
	res, err := DoRequest("DELETE",url,dataJ)
	if err != nil {
		t.Fatal(err)
	}
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}
	fmt.Println("DELETE",path,"=>",string(body))
	if res.StatusCode != want {
		e := fmt.Sprintf("%s request StatusCode want %d but get %d", path, want, res.StatusCode)
		t.Fatal(e)
	}
	return string(body)
}

func DoRequest(method string,url string,data []byte) (resp *http.Response, err error){
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return TestClient.Do(req)
}


func Login_Admin(t *testing.T){
	d := map[string]string{
		"username":"admin",
		"password":"123456",
		"key": "bfbc7f6d0cc10cad0a8ef095d920d3de745f27a4",
	}
	Post_AssertCode(t,"/v1/user/login/test",d,http.StatusOK)
}
