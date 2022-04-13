package main

import (
	"fmt"
	"strings"
)

func trim() {
	str1 := "**This is gaurav Singh***     "
	fmt.Printf("\nsting before  trim %s", str1)

	//res1 := strings.Trim(str1, "gau")
	res2 := strings.TrimSpace(str1)

	fmt.Printf("\nsting after trim %s", res2)
}

func main() {
	trim()
}
