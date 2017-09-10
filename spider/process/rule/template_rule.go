package rule

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"log"
	"strings"
)

func TemplateProcess(rule map[string]string,htmlBytes []byte) interface{}{

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlBytes))
	if err != nil {
		log.Fatal(err)
	}

	resultType := "map"
	rootSel := ""

	v,ok := rule["node"]
	if ok{
		contentInfo := strings.Split(v,"|")
		resultType = contentInfo[0]
		rootSel = contentInfo[1]
	}

	if resultType == "array"{
		result := []interface{}{}
		doc.Find(rootSel).Each(func(i int, s *goquery.Selection) {
			data := getMapFromDom(rule,s)
			result = append(result,data)
		})
		return result
	}

	if resultType == "map"{
		return getMapFromDom(rule,doc.Selection)
	}

	return struct{}{}
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
