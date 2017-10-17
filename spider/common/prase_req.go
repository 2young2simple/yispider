package common

import (
	"YiSpider/spider/model"
	"strings"
	"strconv"
	"fmt"
)

func PraseReq(reqs []*model.Request,ctx map[string]interface{}) []*model.Request{
	resultsReqs := []*model.Request{}
	for _,req := range reqs{
		results,ok := isRuleReq(req,ctx)
		if ok{
			resultsReqs = append(resultsReqs,results...)
		}else{
			resultsReqs = append(resultsReqs,req)
		}
	}
	return resultsReqs
}

func isRuleReq(req *model.Request,ctx map[string]interface{}) ([]*model.Request,bool){
	reqs := []*model.Request{}
	isMatch := false

	findRule := false
	rule := ""
	for i := 0;i< len(req.Url);i++{
		if findRule{
			if req.Url[i] == '}'{
				findRule = false
				isMatch = true
				break
			}
			rule = rule + string(req.Url[i])
		}

		if req.Url[i] == '{'{
			findRule = true
		}
	}

	if !isMatch {
		return nil,false
	}

	if ctx != nil{
		reqs,isMatch = PraseParamCtx(req,rule,ctx)
		if isMatch{
			return reqs,true
		}
	}

	reqs,isMatch = PraseOffset(req,rule)
	if isMatch{
		return reqs,true
	}
	reqs,isMatch = PraseOr(req,rule)
	if isMatch{
		return reqs,true
	}

	return reqs,isMatch
}

// http://xxxxxxxx.com/abc/{begin-end,offset}/   example:{1-400,10}
func PraseOffset(req *model.Request,rule string) ([]*model.Request,bool){
	reqs := []*model.Request{}

	sp := strings.Split(rule,",")
	if len(sp) != 2{
		return reqs,false
	}

	rs := strings.Split(sp[0],"-")
	var begin,end,offset int
	var err error
	begin,err = strconv.Atoi(rs[0])
	end,err = strconv.Atoi(rs[1])
	offset,err = strconv.Atoi(sp[1])
	if err != nil{
		return reqs,false
	}
	if offset == 0{
		return reqs,false
	}
	fmt.Println("begin",begin,"end",end,"offset",offset)
	for i:=begin;i < end;i = i + offset{
		url := strings.Replace(req.Url,"{"+rule+"}",strconv.Itoa(i)  ,1)
		req := &model.Request{Url:url,Method:req.Method,ContentType:req.ContentType,Data:req.Data,Header:req.Header,Cookies:req.Cookies,ProcessName:req.ProcessName}
		reqs = append(reqs,req)
	}

	return reqs,true
}

// http://xxxxxxxx.com/abc/{id1|id2|id3}/
func PraseOr(req *model.Request,rule string) ([]*model.Request,bool){
	reqs := []*model.Request{}

	sp := strings.Split(rule,"|")
	if len(sp) < 2{
		return nil,false
	}

	for _,word := range sp{
		url := strings.Replace(req.Url,"{"+rule+"}",word,1)
		req := &model.Request{Url:url,Method:req.Method,ContentType:req.ContentType,Data:req.Data,Header:req.Header,Cookies:req.Cookies,ProcessName:req.ProcessName}
		reqs = append(reqs,req)
	}

	return reqs,true
}

// http://xxxxxxxx.com/abc/{name}/{id}/
func PraseParamCtx(req *model.Request,rule string,ctx map[string]interface{}) ([]*model.Request,bool){
	reqs := []*model.Request{}

	contains := strings.Contains(rule,"|")
	contains = strings.Contains(rule,",")

	if contains{
		return nil,false
	}

	ruleUrl,_ := ctx[rule]
	urlArray,ok :=  ruleUrl.([]string)
	if ok{
		for _,url := range urlArray{
			u := strings.Replace(req.Url,"{"+rule+"}",string(url),1)
			r := &model.Request{Url:u,Method:req.Method,ContentType:req.ContentType,Data:req.Data,Header:req.Header,Cookies:req.Cookies,ProcessName:req.ProcessName}
			reqs = append(reqs,r)
		}
		return reqs,true
	}

	urlStr,ok := ruleUrl.(string)
	if ok{
		url := strings.Replace(req.Url,"{"+rule+"}",string(urlStr),1)
		r := &model.Request{Url:url,Method:req.Method,ContentType:req.ContentType,Data:req.Data,Header:req.Header,Cookies:req.Cookies,ProcessName:req.ProcessName}
		reqs = append(reqs,r)
		return reqs,true
	}

	return reqs,true
}

