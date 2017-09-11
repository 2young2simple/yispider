package file

import (
	"os"
	"YiSpider/common/model"
	"fmt"
	"time"
	"YiSpider/spider/logger"
	"encoding/json"
)

type FilePipline struct{
	root string
	files map[string]*os.File
}

func NewFilePipline(root string) *FilePipline{
	return &FilePipline{root: root, files: make(map[string]*os.File)}
}

func (c *FilePipline)ProcessData(v interface{},task *model.Task){
	file,ok := c.files[task.Id]
	if !ok{
		var f *os.File
		var err error

		path := fmt.Sprintf("%s%s_result.txt",c.root,task.Name)
		if f,err = os.OpenFile(path,os.O_CREATE|os.O_RDWR,0666);err != nil{
			logger.Error("FilePipline Open File fail, path =",path,err)
			return
		}
		f.WriteString(fmt.Sprintf("========= Task : %s =============\n",task.Name))
		f.WriteString(fmt.Sprintf("======= Task Begin : %s =============\n",time.Now()))

		c.files[task.Id] = f
		file = f
	}
	values,ok :=  v.([]map[string]string)
	if ok{
		for _,value := range values{
			data,err := json.Marshal(value)
			if err != nil{
				logger.Error("FilePipline json.Marshal fail, v = ",v)
				return
			}
			file.WriteString(string(data)+"\n")
		}
	}else{
		data,err := json.Marshal(v)
		if err != nil{
			logger.Error("FilePipline json.Marshal fail, v = ",v)
			return
		}
		file.WriteString(string(data)+"\n")
	}
	logger.Info("File Pipline write.")
	return
}

func (c *FilePipline) Close(){
	for _,f := range c.files{
		f.Close()
	}
}