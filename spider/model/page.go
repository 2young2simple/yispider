package model

type Page struct {
	Result      []interface{}
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

func (p *Page) AddResult(value interface{}) {
	if p.Result == nil {
		p.Result = []interface{}{}
	}
	p.Result = append(p.Result, value)
	p.ResultCount++
}
