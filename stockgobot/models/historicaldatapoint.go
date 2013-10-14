package models

type HistoricalDataPoint struct {
	Symbol string
	Year   int
	Month  int
	Day    int
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}
