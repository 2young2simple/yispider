package rule

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"strings"
	"YiSpider/spider/logger"
	"YiSpider/spider/process/filter"
	"YiSpider/spider/model"
)

func TemplateProcess(task *model.Task,htmlBytes []byte) *model.Page{

	rule := task.Process.TemplateRule.Rule

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
	if err != nil {
		logger.Error("NewDocumentFromReader fail,",err)
	}

	urls := []string{}

	doc.Find("a").Each(func (i int,sel *goquery.Selection){
		href,_ := sel.Attr("href")
		if filter.Filter(href,task){
			urls = append(urls,href)
		}
	})


	resultType := "map"
	rootSel := ""
	page := &model.Page{Urls:urls}

	v,ok := rule["node"]
	if ok{
		contentInfo := strings.Split(v,"|")
		resultType = contentInfo[0]
		rootSel = contentInfo[1]
	}

	if resultType == "array"{
		result := []map[string]string{}
		doc.Find(rootSel).Each(func(i int, s *goquery.Selection) {
			data := getMapFromDom(rule,s)
			result = append(result,data)
		})
		page.Result = result
	}

	if resultType == "map"{
		page.Result = getMapFromDom(rule,doc.Selection)
	}

	return page
}

func getMapFromDom(rule map[string]string,node *goquery.Selection) map[string]string{
	result := make(map[string]string)

	for key,value := range rule{

		if key == "node"{
			continue
		}

		rules := strings.Split(value,"|")
		ValueType := strings.Split(rules[0],".")


		if len(rules) < 2{
			continue
		}

		s := node.Find(rules[1])
		switch ValueType[0] {
			case "text":
				result[key] = s.Text()
			case "html":
				result[key],_ = s.Html()
			case "attr":
				if len(ValueType) < 2{
					continue
				}
				result[key],_ = s.Attr(ValueType[1])
			default:
				result[key] = " "
		}

	}
	return result

}
