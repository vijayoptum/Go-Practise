package main

import "fmt"

func main() {
	fmt.Println("main block starts here :")
	fmt.Println(reversestr("vijay"))
}

func reversestr(s string) string {
	slc := []rune(s)

	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]

	}
	return string(slc)
}
