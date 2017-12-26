package console

import (
	"encoding/json"
	"fmt"
)

type ConsolePipline struct {
}

func NewConsolePipline() *ConsolePipline {
	return &ConsolePipline{}
}

func (c *ConsolePipline) ProcessData(v []map[string]interface{}, taskName string, processName string) {
	bytes, _ := json.Marshal(v)
	fmt.Println("Pipline :", string(bytes))
}
