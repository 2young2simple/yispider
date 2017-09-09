package model

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url string  `json:"url"`
	Method string `json:"method"`

	Header map[string]string `json:"header"`
	Cookies Cookies `json:"cookies"`
	Proxys []string `json:"proxys"`

	RequestBody RequestBody `json:"request_body"`

	Depth int   `json:"depth"`
	EndCount int   `json:"end_count"`
}

type RequestBody struct {
	Type string `json:"type"`  // json urlencode form
	Data map[string]string `json:"data"`
}

type Cookies struct {
	Url string `json:"url"`
	Data string `json:"data"`
}

