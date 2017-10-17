package model

type Page struct {
	Result interface{}
	ResultCount int
	Urls  []*Request
}
