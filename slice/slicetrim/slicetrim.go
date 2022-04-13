package main

import (
	"bytes"
	"fmt"
)

func test(xyz string) {
	slice := []byte(xyz)

	fmt.Printf("\n slice value %d", slice)
	slice[1] = 103
	fmt.Printf("\n slice value %d%s", slice, slice)
	//fmt.Printf("\n sort slice value %s", slice1)

}

func main() {
	//slice := []byte{'g', 'a', 'u', 'r', 'a', 'v', '!', '&'}

	slice := []byte("My name is Gaurav Singh")
	fmt.Printf("\nslice before trim is %s", slice)

	slicechange := bytes.Trim(slice, "Sin")
	fmt.Printf("\nslice after trim is %s", (slicechange))

	abc := "gaurav"

	test(abc)

}

// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func main() {

// 	// Creating and trimming
// 	// the slice of bytes
// 	// Using Trim function
// 	res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
// 	res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
// 	res3 := bytes.Trim([]byte("^^Geek&&"), "$")

// 	// Display the results
// 	fmt.Printf("Final Slice:\n")
// 	fmt.Printf("\nSlice 1: %s", res1)
// 	fmt.Printf("\nSlice 2: %s", res2)
// 	fmt.Printf("\nSlice 3: %s", res3)
// }
