package importing

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/julianduniec/news/stockgobot/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ImportHistory(symbol string) ([]*models.HistoricalDataPoint, error) {
	resp, err := http.Get(getUrl(symbol))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return nil, errors.New("Could not fetch data")
	}
	data := make([]*models.HistoricalDataPoint, 0)
	reader := bufio.NewReader(resp.Body)
	count := 0
	for {
		s, err := reader.ReadString(10)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if count > 0 {
			data = append(data, createHistoricalDataPointFromRow(symbol, s))
		}
		count++
	}
	return data, nil
}

func createHistoricalDataPointFromRow(symbol, row string) *models.HistoricalDataPoint {
	split := strings.Split(row, ",")
	date := parseTime(split[0])
	open, _ := strconv.ParseFloat(split[1], 32)
	high, _ := strconv.ParseFloat(split[2], 32)
	low, _ := strconv.ParseFloat(split[3], 32)
	closeRate, _ := strconv.ParseFloat(split[4], 32)
	volume, _ := strconv.Atoi(split[5])
	return &models.HistoricalDataPoint{symbol, date, open, high, low, closeRate, volume}
}

func parseTime(s string) time.Time {
	split := strings.Split(s, "-")
	year, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func getUrl(symbol string) string {
	return fmt.Sprintf("http://ichart.finance.yahoo.com/table.csv?s=%s&e=%00d&f=%d&g=d&ignore=.csv",
		symbol,
		int(time.Now().Month()), time.Now().Year())
}
