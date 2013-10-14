package importing

import (
	"bufio"
	"github.com/julianduniec/news/stockgobot/models"
	"io"
	"os"
	"strings"
)

func ImportSymbols() []*models.Symbol {
	file, err := os.Open("./data/nasdaqlisted.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	symbols := make([]*models.Symbol, 0)
	for {
		s, err := reader.ReadString(10)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		symbol := symbolFromRow(s)
		if symbol != nil {
			symbols = append(symbols, symbol)
		}
	}
	return symbols
}

func symbolFromRow(row string) *models.Symbol {
	split := strings.Split(row, "|")
	symbol := split[0]
	name := split[1]
	if len(name) != 0 && len(symbol) != 0 {
		return &models.Symbol{symbol, name}
	}
	return nil
}
