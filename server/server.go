package server

import (
	"code.google.com/p/gorest"
	"fmt"
	"net/http"
)

func Start() {
	gorest.RegisterService(new(NewsService))
	http.Handle("/",gorest.Handle())    
	http.ListenAndServe(":8080",nil)
}

type News struct {
	Title			string
	Preamble 		string
	Body 			string
}

type NewsService struct {
	gorest.RestService						`consumes:"application/json" produces:"application/json"`
	listNews				gorest.EndPoint `method:"GET" path:"/news" output:"News"`
	numbers					gorest.EndPoint `method:"GET" path:"/numbers" output:"[]News"`
}

func(s NewsService) Numbers() []News {
	return []News{
		News{"Hello", "World", "Of news"}}
}

func(s NewsService) ListNews() News {
	return News{"Hello", "World", "Of news"}
}