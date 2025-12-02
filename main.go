package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
) 

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for ;; {
		fmt.Print("Pokedex > ")

		ok:=scanner.Scan()
		if !ok {
			break
		}

		input:=strings.ToLower(scanner.Text())

		if cmd, ok := commands[input]; ok {
			cmd.callback()
		}
		

	}

	return
}
