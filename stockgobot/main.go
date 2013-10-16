package main

import (
	"bufio"
	"fmt"
	"github.com/julianduniec/news/stockgobot/importing"
	"github.com/julianduniec/news/stockgobot/store"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	res, _ := reader.ReadString(10)
	return res[0 : len(res)-1]
}

func main() {

	printWelcomeMessage()
	fmt.Print(" > ")
	for s := readLine(); s != "quit"; s = readLine() {
		switch s {
		case "help":
			fallthrough
		case "man":
			fallthrough
		case "menu":
			printMenu()
		case "symbols":
			printSymbols()
		case "import":
			importing.Run()
			fmt.Println("Done")
		}
		fmt.Print(" > ")
	}
}

func printMenu() {
	fmt.Println("help       -> prints this manual")
	fmt.Println("symbols    -> list symbols")
	fmt.Println("import     -> execute import")
	fmt.Println("quit       -> quit")
	fmt.Println("")
}

func printWelcomeMessage() {
	fmt.Println("")
	fmt.Println("*************************************************************")
	fmt.Println("* Welcome to Stock-Go-Bot. What would you like to do today? *")
	fmt.Println("*************************************************************")
	fmt.Println("")
}

func printSymbols() {
	symbols := store.GetSymbols()
	for i, symbol := range symbols {
		fmt.Println(symbol.Symbol)
		if i != 0 && i%10 == 0 {
			fmt.Println("Type 'it' for more, or enter to go back")
			fmt.Print(" > ")
			if readLine() != "it" {
				break
			}
		}
	}
}
