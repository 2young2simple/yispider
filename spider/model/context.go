package model

import "net/http"

type Context struct {
	Body []byte
	Request *http.Request
	Header http.Header
}

func (c *Context) Clear(){
	c.Body = nil
	c.Request = nil
	c.Header = nil
}