package store


type News struct {
	Url				string  		`json:"url"`
	Title			string			`json:"title"`
	Preamble 		string 			`json:"preamble"`
	Body 			string			`json:"body"`
}

var (
	news = []News{}
)

func All() []News {
	return news
}

func Add(newsItem News) {
	news = append(news, newsItem)
}