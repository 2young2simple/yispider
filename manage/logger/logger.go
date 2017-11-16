package logger

import "fmt"

func Info(v ...interface{}) {
	fmt.Println(v)
}

func Debug(v ...interface{}) {
	fmt.Println(v)
}

func Warn(v ...interface{}) {
	fmt.Println(v)
}

func Error(v ...interface{}) {
	fmt.Println(v)
}
