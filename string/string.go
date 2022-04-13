package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	mystr := "Welcome**"
	bytemystr := "Welcome**99999"
	fmt.Printf("String: %s", mystr)

	mystr1 := []rune(mystr)
	mystr2 := []byte(bytemystr)

	for i, j := range mystr1 {
		fmt.Printf("\nindex is %c is %v", j, j)
		mystr1[i] = 97
		println(mystr[i], j)

	}
	countabc := string(mystr1)
	fmt.Printf("String: %s", countabc)

	fmt.Printf("\nbyte string %v", mystr2)
	fmt.Printf("\nrune string %v", mystr1)

	count := len(mystr2)
	println(count)
	count1 := utf8.RuneCountInString(string(mystr1))
	println(count1)

}
