package rule

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"encoding/json"
)

func TestTemplateProcess(t *testing.T) {
	doc,err := goquery.NewDocument("https://www.qiushibaike.com/")
	if err != nil{
		t.Fatal("open url fail ",err)
	}
	html,err := doc.Html()
	if err != nil{
		t.Fatal("get html fail ",err)
	}
	rule := map[string]string{
		"node":"array|.article",
		"url":"attr.href|.contentHerf",
		"author":"attr.alt|.author a img",
		"content":"text|.content span",
		"like_num":"text|.stats-vote i",
		"comment_num":"text|.stats-comments a i",
	}

	result := TemplateProcess(rule,[]byte(html))
	data,_ := json.Marshal(result)
	fmt.Println("Result :",string(data))
}