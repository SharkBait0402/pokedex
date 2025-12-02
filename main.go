package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
) 

func commandExit() error {
	fmt.Println("Closing the pokedex... Goodbye!")
	os.Exit(0)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for ;; {
		fmt.Print("Pokedex > ")

		ok:=scanner.Scan()
		if !ok {
			break
		}

		line:=scanner.Text()
		

	}

	return
}
