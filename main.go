package main

import(
	"github.com/julianduniec/news/server"
	"github.com/julianduniec/news/importer"
	"flag"
)



func main() {
	//Load flags
	var port string
	flag.StringVar(&port, "port", "8090", "Port number")
	routineQuit := make(chan int)

	go importer.Start()

	go server.Start(port)

	<-routineQuit
}