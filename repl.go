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

func commandWord(text string) string {
	txt := strings.ToLower(text)
	txt=strings.TrimSpace(txt)
	words:=strings.Fields(txt)
	return words[0]
}

func restInput(text string) string {
 txt:= strings.ToLower(text)
 txt = strings.TrimSpace(txt)
 words:=strings.Fields(txt)

 final:=""

 for i:=1;i<len(words);i++ {
	final += words[i] + "-"
 }

 return final
}
