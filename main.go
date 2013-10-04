package main

import(
	"github.com/julianduniec/news/server"
	"github.com/julianduniec/news/importer"
	"flag"
	"time"
)

var(
	port 				string
	pollingfrequency 	string
)

func init() {
	flag.StringVar(&port, 				"port", 				"8090", "Port number")
	flag.StringVar(&pollingfrequency, 	"pollingfrequency", 	"2m0s", "Polling frequency")
}

func main() {
	flag.Parse()
	
	duration, _ := time.ParseDuration(pollingfrequency)

	go importer.Start(duration)

	go server.Start(port)

	//RUN FOREVER muahahaha
	<-make(chan int)
}