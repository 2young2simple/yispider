package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	"YiSpider/spider/pipline/file"
	spider2 "YiSpider/spider/spider"
	"encoding/json"
)

type Movies struct {
	Datas []Movie `json:"data"`
}
type Movie struct {
	Rate  string   `json:"rate"`
	Start string   `json:"start"`
	Id    string   `json:"id"`
	Url   string   `json:"url"`
	Title string   `json:"title"`
	Cover string   `json:"cover"`
	Casts []string `json:"casts"`
}

type PageProcess struct{}

func (p *PageProcess) Process(context model.Context) (*model.Page, error) {
	movies := Movies{}
	if err := json.Unmarshal(context.Body, &movies); err != nil {
		return nil, err
	}
	page := &model.Page{}
	for _, movie := range movies.Datas {
		page.AddResult(movie)
	}
	return page, nil
}

func main() {
	sp := &spider2.Spider{}
	sp.Name = "douban-movie-code"
	sp.Id = "douban-movie-code"
	sp.Requests = []*model.Request{
		{
			Method:      "get",
			Url:         "https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start={0-10000,20}",
			ProcessName: "movie",
		},
	}
	sp.AddProcess("movie", &PageProcess{})
	sp.Pipline = file.NewFilePipline("./")

	app := spider.New()
	app.AddSpider(sp)
	app.Run()
}
