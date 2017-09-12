package discover

type Node struct {
	IsHealth bool
	IP       string
	Name     string
	CPU      int
	MetaData map[string]string
}

type WorkerInfo struct {
	Name string		`json:"name"`
	IP   string		`json:"ip"`
	CPU  int		`json:"cpu"`
	MetaData map[string]string `json:"metadata"`
}
