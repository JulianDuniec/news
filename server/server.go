package server

import (
	"github.com/julianduniec/gorest"
	"github.com/julianduniec/news/store"
	"net/http"
)

func Start(port string) {
	gorest.RegisterService(new(NewsService))
	http.Handle("/", gorest.Handle())    
	http.ListenAndServe(":" + port, nil)
}


type NewsService struct {
	gorest.RestService						`consumes:"application/json" produces:"application/json"`
	listNews				gorest.EndPoint `method:"GET" path:"/news/{from:int}/{to:int}" output:"[]News"`
	addNews					gorest.EndPoint	`method:"PUT" path:"/news" postdata:"News"`
}

func(s NewsService) ListNews(from, to int) []store.News {
	return store.All(from, to)
}

func(s NewsService) AddNews(news store.News) {
	store.Add(news)
	s.ResponseBuilder().Created("/news")
}