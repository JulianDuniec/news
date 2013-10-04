package server

import (
	"code.google.com/p/gorest"
	"github.com/julianduniec/news/store"
	"net/http"
	
)

func Start() {
	gorest.RegisterService(new(NewsService))
	http.Handle("/",gorest.Handle())    
	http.ListenAndServe(":8080",nil)
}


type NewsService struct {
	gorest.RestService						`consumes:"application/json" produces:"application/json"`
	listNews				gorest.EndPoint `method:"GET" path:"/news" output:"[]News"`
	addNews					gorest.EndPoint	`method:"PUT" path:"/news" postdata:"News"`
}

func(s NewsService) ListNews() []store.News {
	return store.All()
}

func(s NewsService) AddNews(news store.News) {
	store.Add(news)
	s.ResponseBuilder().Created("/news")
}