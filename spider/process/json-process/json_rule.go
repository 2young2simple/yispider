package json_process

import (
	"YiSpider/spider/model"
	"YiSpider/spider/logger"
	"strings"
	simplejson "github.com/bitly/go-simplejson"
)

func JsonRuleProcess(process *model.Process,bytes []byte) (*model.Page,error){
	return Process(process.JsonRule.Rule,bytes)
}

func Process(jsonRule map[string]string,bytes []byte) (*model.Page,error){
	page := &model.Page{}

	sJson,err := simplejson.NewJson(bytes)
	if err != nil {
		logger.Error("NewDocumentFromReader fail,",err)
		return nil,err
	}

	resultType := "map"
	rootSel := []string{}

	v,ok := jsonRule["node"]

	if ok{
		contentInfo := strings.Split(v,"|")
		resultType = contentInfo[0]
		selStr := contentInfo[1]
		rootSel = strings.Split(selStr,".")
	}

	if resultType == "array"{
		result := []map[string]interface{}{}

		for _,name := range rootSel{
			sJson = sJson.Get(name)
		}
		rootNode,err := sJson.Array()
		if err != nil {
			logger.Error("Json fail,",err)
			return nil,err
		}
		if len(rootNode) >= 0{
			for _,node := range rootNode{
				nodeMap,ok := node.(map[string]interface{})
				if !ok{
					continue
				}
				data := map[string]interface{}{}
				for key,value := range jsonRule{
					if key == "node"{
						continue
					}
					data[key] = nodeMap[value]
				}
				result = append(result,data)
			}

		}

		page.Urls = []*model.Request{}
		page.Result = result
		page.ResultCount = len(result)
	}

	if resultType == "map"{

		result := map[string]interface{}{}

		for _,name := range rootSel{
			sJson = sJson.Get(name)
		}
		rootNode,err := sJson.Map()
		if err != nil {
			logger.Error("Json fail,",err)
			return nil,err
		}

		for key,value := range jsonRule{
			result[key] = rootNode[value]
		}

		page.Urls = []*model.Request{}
		page.Result = result
		page.ResultCount = 1
	}
	return page,nil
}