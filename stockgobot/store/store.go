package store

import (
	"github.com/julianduniec/news/stockgobot/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	session *mgo.Session
)

func Init() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
}
func SaveSymbol(symbol *models.Symbol) {
	if session == nil {
		Init()
	}
	session.DB("stock").C("symbols").Upsert(symbol, symbol)
}

func ClearHistory(symbol string) {
	if session == nil {
		Init()
	}
	session.DB("stock").C("history").RemoveAll(bson.M{"symbol": symbol})
}

func SaveHistory(history *models.HistoricalDataPoint) {
	if session == nil {
		Init()
	}
	session.DB("stock").C("history").Insert(history)
}

func GetSymbols() []models.Symbol {
	if session == nil {
		Init()
	}
	result := []models.Symbol{}
	session.DB("stock").C("symbols").Find(bson.M{}).All(&result)
	return result
}
