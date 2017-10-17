package model


type Task struct {
	Id   string `json:"id"`
	Name string `jsonTask:"name"`

	Request []*Request `json:"request"`
	Process []Process `json:"process"`
	Pipline string `json:"pipline"`

	Depth int   `json:"depth"`
	EndCount int   `json:"end_count"`

}

type Request struct {
	Url string `json:"url"`
	Method string `json:"method"`
	ContentType string `json:"type"`  // json urlencode form
	Data map[string]string `json:"data"`
	Header map[string]string `json:"header"`
	Cookies Cookies `json:"cookies"`
	ProcessName string `json:"process_name"`
}

type Cookies struct {
	Url string `json:"url"`
	Data string `json:"data"`
}

type Process struct {
	Name string		`json:"name"`
	RegUrl []string	`json:"reg_url"`
	Type string `json:"type"`  // template json self_process
	TemplateRule TemplateRule `json:"template_rule"`
	JsonRule JsonRule `json:"json_rule"`
	AddQueue []*Request `json:"add_queue"`   //  http://www.baidu.com/{name}/{ctx}
}

type TemplateRule struct {
	Rule map[string]string
}

type JsonRule struct {
	Rule map[string]string
}
