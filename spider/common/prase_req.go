package common

import (
	"YiSpider/spider/model"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

func PraseReq(reqs []*model.Request, ctx map[string]interface{}) []*model.Request {
	resultsReqs := []*model.Request{}
	for _, req := range reqs {
		results, ok := isRuleReq(req, ctx)
		if ok {
			resultsReqs = append(resultsReqs, results...)
		} else {
			resultsReqs = append(resultsReqs, req)
		}
	}
	return resultsReqs
}

func FindRule(text string) [][]string {
	reg := regexp.MustCompile(`{([^}]+)}`)
	return reg.FindAllStringSubmatch(text, -1)
}

func isRuleReq(req *model.Request, ctx map[string]interface{}) ([]*model.Request, bool) {
	reqs := []*model.Request{req}
	outReqs := []*model.Request{}
	finalReqs := []*model.Request{}
	isMatch := false

	rules := FindRule(req.Url)
	if len(rules) > 0 {
		isMatch = true
	} else {
		return nil, false
	}

	if ctx != nil {
		reqs, isMatch = PraseParamCtx(req, rules, ctx)
	}
	for _, r := range reqs {
		outReqs = append(outReqs, PraseOffset(r)...)
	}

	for _, r := range outReqs {
		finalReqs = append(finalReqs, PraseOr(r)...)
	}

	if isMatch {
		return finalReqs, true
	}

	return finalReqs, isMatch
}

// http://xxxxxxxx.com/abc/{begin-end,offset}/   example:{1-400,10}
func PraseOffset(req *model.Request) []*model.Request {
	reqs := []*model.Request{}
	outrReqs := []*model.Request{}

	rules := FindRule(req.Url)
	if len(rules) <= 0 {
		return []*model.Request{req}
	}

	rule := rules[0][1]
	sp := strings.Split(rule, ",")

	if len(sp) != 2 {
		return []*model.Request{req}
	}

	rs := strings.Split(sp[0], "-")
	var begin, end, offset int
	var err error
	begin, err = strconv.Atoi(rs[0])
	end, err = strconv.Atoi(rs[1])
	offset, err = strconv.Atoi(sp[1])
	if err != nil {
		return []*model.Request{req}
	}
	if offset == 0 {
		return []*model.Request{req}
	}

	for i := begin; i <= end; i = i + offset {
		url := strings.Replace(req.Url, "{"+rule+"}", strconv.Itoa(i), 1)
		req := &model.Request{Url: url, Method: req.Method, ContentType: req.ContentType, Data: req.Data, Header: req.Header, Cookies: req.Cookies, ProcessName: req.ProcessName}
		reqs = append(reqs, req)
	}

	for _, r := range reqs {
		outrReqs = append(outrReqs, PraseOffset(r)...)
	}

	return outrReqs
}

// http://xxxxxxxx.com/abc/{id1|id2|id3}/
func PraseOr(req *model.Request) []*model.Request {
	reqs := []*model.Request{}
	outrReqs := []*model.Request{}

	rules := FindRule(req.Url)
	if len(rules) <= 0 {
		return []*model.Request{req}
	}
	ruleArray := rules[0]
	rule := ruleArray[1]
	sp := strings.Split(rule, "|")
	if len(sp) < 2 {
		return []*model.Request{req}
	}

	for _, word := range sp {
		url := strings.Replace(req.Url, "{"+rule+"}", word, 1)
		r := &model.Request{Url: url, Method: req.Method, ContentType: req.ContentType, Data: req.Data, Header: req.Header, Cookies: req.Cookies, ProcessName: req.ProcessName}
		reqs = append(reqs, r)
	}

	for _, r := range reqs {
		outrReqs = append(outrReqs, PraseOr(r)...)
	}

	return outrReqs
}

// http://xxxxxxxx.com/abc/{name}/{id}/
func PraseParamCtx(req *model.Request, rules [][]string, ctx map[string]interface{}) ([]*model.Request, bool) {
	reqs := []*model.Request{}
	reqUrl := req.Url

	count := strings.Count(reqUrl, "$")
	if count <= 0 {
		return []*model.Request{req}, false
	}

	for ctxName, ruleUrl := range ctx {
		urlArray, ok := ruleUrl.([]string)
		if ok {
			for _, url := range urlArray {
				u := strings.Replace(reqUrl, "{$"+url+"}", string(url), -1)
				u = strings.Replace(reqUrl, "$"+url, string(url), -1)
				r := &model.Request{Url: u, Method: req.Method, ContentType: req.ContentType, Data: req.Data, Header: req.Header, Cookies: req.Cookies, ProcessName: req.ProcessName}
				if newCount := strings.Count(u, "$"); newCount != count {
					reqUrl = u
					count = newCount
					if count == 0 {
						reqs = append(reqs, r)
					}
				}
			}
		}

		urlStr, ok := ruleUrl.(string)
		if ok {
			url := strings.Replace(reqUrl, "{$"+ctxName+"}", string(urlStr), -1)
			url = strings.Replace(url, "$"+ctxName, string(urlStr), -1)
			r := &model.Request{Url: url, Method: req.Method, ContentType: req.ContentType, Data: req.Data, Header: req.Header, Cookies: req.Cookies, ProcessName: req.ProcessName}
			if newCount := strings.Count(url, "$"); newCount != count {
				reqUrl = url
				count = newCount
				if count == 0 {
					reqs = append(reqs, r)
				}
			}
		}

		urlNumber, ok := ruleUrl.(json.Number)
		if ok {
			url := strings.Replace(reqUrl, "{$"+ctxName+"}", string(urlNumber), -1)
			url = strings.Replace(url, "$"+ctxName, string(urlNumber), -1)
			r := &model.Request{Url: url, Method: req.Method, ContentType: req.ContentType, Data: req.Data, Header: req.Header, Cookies: req.Cookies, ProcessName: req.ProcessName}

			if newCount := strings.Count(url, "$"); newCount != count {
				reqUrl = url
				count = newCount
				if count == 0 {
					reqs = append(reqs, r)
				}
			}
		}

		urlInt, ok := ruleUrl.(int)
		if ok {
			url := strings.Replace(reqUrl, "{$"+ctxName+"}", strconv.Itoa(urlInt), -1)
			url = strings.Replace(url, "$"+ctxName, strconv.Itoa(urlInt), -1)
			r := &model.Request{Url: url, Method: req.Method, ContentType: req.ContentType, Data: req.Data, Header: req.Header, Cookies: req.Cookies, ProcessName: req.ProcessName}
			if newCount := strings.Count(url, "$"); newCount != count {
				reqUrl = url
				count = newCount
				if count == 0 {
					reqs = append(reqs, r)
				}
			}
		}

	}

	return reqs, true
}
