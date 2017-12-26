package pipline

type Pipline interface {
	ProcessData(v []map[string]interface{}, taskName string, processName string)
}
