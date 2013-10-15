package main

import (
	"bufio"
	"fmt"
	"github.com/julianduniec/news/stockgobot/importing"
	"os"
)

func readLine(reader *bufio.Reader) string {
	res, _ := reader.ReadString(10)
	return res[0 : len(res)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Println("************************************************************")
	fmt.Println("* Welcome to Stock-Go-Bot. What would you like to do today?*")
	fmt.Println("************************************************************")
	fmt.Println("")
	for s := readLine(reader); s != "quit"; s = readLine(reader) {
		switch s {
		case "hello":
			fmt.Println("Hello yourself!")
		case "import":
			importing.Run()
			fmt.Println("Done")
		}
	}

}
