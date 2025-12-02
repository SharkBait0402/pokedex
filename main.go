package main

import(
	"fmt"
	"bufio"
	"strings"
) 

func main() {
	scanner := bufio.NewScanner(strings.NewReader(""))

	fmt.Printf("input recieved: %v", scanner.Text)
	return
}
