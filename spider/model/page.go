package model

type Page struct {
	Result      []map[string]interface{}
	ResultCount int
	Urls        []*Request
}

func (p *Page) AddUrl(req *Request) {
	if p.Urls == nil {
		p.Urls = []*Request{}
	}
	p.Urls = append(p.Urls, req)
}

func (p *Page) AddUrls(req []*Request) {
	if p.Urls == nil {
		p.Urls = []*Request{}
	}
	p.Urls = append(p.Urls, req...)
}

func (p *Page) AddResult(value map[string]interface{}) {
	if p.Result == nil {
		p.Result = []map[string]interface{}{}
	}
	p.Result = append(p.Result, value)
	p.ResultCount++
}
