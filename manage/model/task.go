package model

type Task struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Host   string `json:"host"`
	Method string `json:"method"`

	Header  map[string]string `json:"header"`
	Cookies Cookies           `json:"cookies"`
	Proxys  []string          `json:"proxys"`

	RequestBody RequestBody `json:"request_body"`

	Process Process `json:"process"`

	Depth    int `json:"depth"`
	EndCount int `json:"end_count"`

	Pipline string `json:"pipline"`
}

type RequestBody struct {
	Type string            `json:"type"` // json urlencode form
	Data map[string]string `json:"data"`
}

type Cookies struct {
	Url  string `json:"url"`
	Data string `json:"data"`
}

type Process struct {
	Url          string
	RegUrl       []string
	Type         string       `json:"type"` // template json self_process
	TemplateRule TemplateRule `json:"template_rule"`
	JsonRule     JsonRule     `json:"json_rule"`
}

type TemplateRule struct {
	Rule map[string]string
}

type JsonRule struct {
	Rule map[string]interface{}
}
