package main

import(
	"fmt"
) 

func main() {
	words := cleanInput("Hello  World")
	for _, word := range words{
		fmt.Println(word)
	}
}
