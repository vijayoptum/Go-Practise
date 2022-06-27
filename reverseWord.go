package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("main block starts here")

	fmt.Println(reverseword("vijay battula raju"))
}

func reverseword(s string) string {

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")

}
