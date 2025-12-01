package main

import (
	"strings"
)

func cleanInput(text string) []string {
	
	txt := strings.ToLower(text)

	txt = strings.TrimSpace(txt)

	words := strings.Fields(txt)



	return words

}

