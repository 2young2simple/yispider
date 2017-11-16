package pipline

type Pipline interface {
	ProcessData(v []interface{}, taskName string, processName string)
}
