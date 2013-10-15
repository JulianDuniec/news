package importing

import (
	"fmt"
	"github.com/julianduniec/news/stockgobot/models"
	"github.com/julianduniec/news/stockgobot/store"
	"github.com/julianduniec/throttler"
	"sync"
)

var wg sync.WaitGroup

type HistoryImportResult struct {
	result []*models.HistoricalDataPoint
	err    error
}

func toInterfaceSlice(symbols []*models.Symbol) []interface{} {
	res := make([]interface{}, len(symbols))
	for i, s := range symbols {
		res[i] = s
	}
	return res
}

func importSymbolHistoryWorker(in interface{}, c chan interface{}) {
	fmt.Println("Fetching item")
	symbol := in.(*models.Symbol)
	data, err := ImportHistory(symbol.Symbol)
	if err != nil {
		fmt.Println(err)
	}
	c <- HistoryImportResult{data, err}

}
func importSymbolHistoryWorkerComplete(in interface{}) {
	fmt.Println("Finished item")
	result := in.(HistoryImportResult)
	if result.err == nil {
		wg.Add(1)
		go saveHistory(result.result, &wg)
	}
}
func Run() {
	symbols := ImportSymbols()

	saveSymbols(symbols)

	t := throttler.Throttler{
		200,
		importSymbolHistoryWorker,
		importSymbolHistoryWorkerComplete}

	t.Run(toInterfaceSlice(symbols))
	wg.Wait()
	fmt.Println("Done")
}

func saveSymbols(symbols []*models.Symbol) {
	fmt.Println("Saving symbols	")
	for _, s := range symbols {
		store.SaveSymbol(s)
	}

}

func saveHistory(data []*models.HistoricalDataPoint, wg *sync.WaitGroup) {
	store.ClearHistory(data[0].Symbol)
	for _, d := range data {
		store.SaveHistory(d)
	}
	wg.Done()
}
