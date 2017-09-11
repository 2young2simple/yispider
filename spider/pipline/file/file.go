package file

import "fmt"

type ConsolePipline struct{

}

func NewConsolePipline() *ConsolePipline{
	return &ConsolePipline{}
}

func (c *ConsolePipline)ProcessData(v interface{}){
	fmt.Println("Pipline :",v)
}