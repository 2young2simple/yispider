package discover

type WorkerInfo struct {
	Name string		`json:"name"`
	IP   string		`json:"ip"`
	CPU  int		`json:"cpu"`
	MetaData map[string]string `json:"metadata"`
}
