package utils

import (
	"os/exec"
	"os"
	//"strings"
	"fmt"
)

func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	//i := strings.LastIndex(s, "\\")
	//path := string(s[0 : i+1])
	fmt.Println("path",s)
	return s
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}