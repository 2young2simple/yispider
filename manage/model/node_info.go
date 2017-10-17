package model

type Node struct {
	IsHealth bool						`json:"is_health"`
	IP       string						`json:"ip"`
	Name     string						`json:"name"`
	CPU      int						`json:"cpu"`
	MetaData map[string]string			`json:"metadata"`
	SpiderData map[string]*SpiderData	`json:"spider_data"`
}

type WorkerInfo struct {
	Name string		`json:"name"`
	IP   string		`json:"ip"`
	CPU  int		`json:"cpu"`
	MetaData map[string]string `json:"metadata"`
	SpiderData map[string]*SpiderData `json:"spider_data"`
}

type SpiderData struct {
	DownloadFailCount int32 `json:"download_fail_count"`
	DownloadCount int32 `json:"download_count"`
	UrlNum int32 `json:"url_num"`
	WaitUrlNum int `json:"wait_url_num"`
	CrawlerResultNum int32 `json:"crawler_result_num"`
}
