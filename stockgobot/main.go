package main

import (
	"fmt"
	"github.com/julianduniec/news/stockgobot/importing"
	"github.com/julianduniec/news/stockgobot/models"
	"github.com/julianduniec/news/stockgobot/store"
	"sync"
)

type HistoryImportResult struct {
	result []*models.HistoricalDataPoint
	err    error
}

func importChunk(symbols []*models.Symbol) []*models.Symbol {
	maxLen := 200
	if len(symbols) < maxLen {
		maxLen = len(symbols)
	}
	fmt.Println("Importing chunk", maxLen, len(symbols))
	resultChannel := make(chan HistoryImportResult)
	for i := 0; i < maxLen; i++ {
		go importHistory(symbols[i].Symbol, resultChannel)
	}
	fmt.Println("Synchronizing")
	var wg sync.WaitGroup
	for i := 0; i < maxLen; i++ {
		currentResult := <-resultChannel
		if currentResult.err == nil {
			wg.Add(1)
			go saveHistory(currentResult.result, &wg)
		}
	}
	fmt.Println("Waiting for all saves")
	wg.Done()
	if len(symbols) > maxLen {
		return symbols[maxLen:len(symbols)]
	}
	return nil

}

func main() {
	store.Init()
	symbols := importing.ImportSymbols()

	fmt.Println("Storing symbols")
	for _, s := range symbols {
		go store.SaveSymbol(s)
	}
	for {
		symbols = importChunk(symbols)
		if symbols == nil {
			break
		}
	}
	fmt.Println("Done")
}

func saveHistory(data []*models.HistoricalDataPoint, wg *sync.WaitGroup) {
	store.SaveHistory(data)
	wg.Done()
}

func importHistory(symbol string, res chan HistoryImportResult) {
	data, err := importing.ImportHistory(symbol)
	if err != nil {
		fmt.Println(err)
	}
	res <- HistoryImportResult{data, err}
}
