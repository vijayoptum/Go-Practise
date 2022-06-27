package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("main block starts here")
	var mainstr, srchstr string
	fmt.Println("\n Enter the main string")
	fmt.Scan(&mainstr)
	fmt.Println("\n Enter the search string")
	fmt.Scan(&srchstr)
	outpt := strings.Contains(mainstr, srchstr)
	fmt.Println("Final output is : ", outpt)
	//fmt.Println(strings.Contains("Vijay battula raju", "raju"))
}
