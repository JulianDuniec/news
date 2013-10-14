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
	session.DB("stock").C("history").Upsert(symbol, symbol)
}

func SaveHistory(history []*models.HistoricalDataPoint) {
	symbol := history[0].Symbol
	session.DB("stock").C("history").RemoveAll(bson.M{"symbol": symbol})
	session.DB("stock").C("history").Insert(history)
}
