package main

import(
	"github.com/julianduniec/news/server"
	"github.com/julianduniec/news/importer"
	"github.com/julianduniec/news/jsonpproxy"
	"flag"
	"time"
)

var(
	port 				string
	proxyport			string
	pollingfrequency 	string
	rssFile				string
)

func init() {
	flag.StringVar(&port, 				"port", 				"8080", "Port number")
	flag.StringVar(&proxyport, 			"proxyport", 			"8090", "Port number")
	flag.StringVar(&pollingfrequency, 	"pollingfrequency", 	"2m0s", "Polling frequency")
	flag.StringVar(&rssFile, 			"rssFile", 				"", "Polling frequency")
}

func main() {
	flag.Parse()

	duration, _ := time.ParseDuration(pollingfrequency)

	go importer.Start(duration, rssFile)

	go server.Start(port)

	go jsonpproxy.Start(proxyport)
	//RUN FOREVER muahahaha
	<-make(chan int)
}