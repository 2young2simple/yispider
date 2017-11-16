package template_process

import (
	"YiSpider/spider/common"
	"YiSpider/spider/logger"
	"YiSpider/spider/model"
	"YiSpider/spider/process/filter"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	url2 "net/url"
	"strings"
)

func TemplateRuleProcess(process *model.Process, context model.Context) (*model.Page, error) {
	page := &model.Page{}

	rule := process.TemplateRule.Rule

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(context.Body))
	if err != nil {
		logger.Error("NewDocumentFromReader fail,", err)
		return nil, err
	}

	if len(process.RegUrl) > 0 {
		doc.Find("a").Each(func(i int, sel *goquery.Selection) {
			href, _ := sel.Attr("href")
			href = getComplateUrl(context.Request.URL, href)
			if filter.Filter(href, process) {
				page.AddUrl(&model.Request{Url: href, Method: "get"})
			}
		})
	}

	resultType := "map"
	rootSel := ""

	v, ok := rule["node"]
	if ok {
		contentInfo := strings.Split(v, "|")
		resultType = contentInfo[0]
		rootSel = contentInfo[1]
	}

	if resultType == "array" {

		doc.Find(rootSel).Each(func(i int, s *goquery.Selection) {
			data := getMapFromDom(rule, s)
			if data == nil {
				return
			}
			if len(process.AddQueue) > 0 {
				page.AddUrls(common.PraseReq(process.AddQueue, data))
			}
			page.AddResult(data)
		})
	}

	if resultType == "map" {
		data := getMapFromDom(rule, doc.Selection)
		if len(process.AddQueue) > 0 {
			page.AddUrls(common.PraseReq(process.AddQueue, data))
		}
		page.AddResult(data)
	}

	return page, nil
}

func getMapFromDom(rule map[string]string, node *goquery.Selection) map[string]interface{} {

	result := make(map[string]interface{})

	isNull := true

	for key, value := range rule {

		if key == "node" {
			continue
		}

		rules := strings.Split(value, "|")
		ValueType := strings.Split(rules[0], ".")

		if len(rules) < 2 {
			continue
		}

		s := node.Find(rules[1])
		switch ValueType[0] {
		case "text":
			result[key] = s.Text()
		case "html":
			result[key], _ = s.Html()
		case "attr":
			if len(ValueType) < 2 {
				continue
			}
			result[key], _ = s.Attr(ValueType[1])
		case "texts":
			arr := []string{}
			s.Each(func(i int, sel *goquery.Selection) {
				text := sel.Text()
				arr = append(arr, text)
			})
			j, _ := json.Marshal(arr)
			result[key] = string(j)
		case "htmls":
			arr := []string{}
			s.Each(func(i int, sel *goquery.Selection) {
				html, _ := s.Html()
				arr = append(arr, html)
			})
			j, _ := json.Marshal(arr)
			result[key] = string(j)
		case "attrs":
			arr := []string{}
			attr := ""
			s.Each(func(i int, sel *goquery.Selection) {
				if len(ValueType) >= 2 {
					attr, _ = sel.Attr(ValueType[1])
					arr = append(arr, attr)
				}
			})
			result[key] = arr
		default:
			result[key] = ""
		}
		res, ok := result[key].(string)
		if ok || len(res) != 0 {
			isNull = false
		}
	}

	if isNull == true {
		return nil
	}

	return result
}

func getComplateUrl(url *url2.URL, href string) string {

	if strings.HasPrefix(href, "/") {
		newHref := fmt.Sprintf("%s://%s%s", url.Scheme, url.Host, href)
		return newHref
	}

	newHref := fmt.Sprintf("%s://%s/%s", url.Scheme, url.Host, href)
	return newHref
}
