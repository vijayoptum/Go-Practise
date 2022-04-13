package main

import "fmt"

func main() {
	s := "abcdcba"
	abc := []byte(s)
	fmt.Printf("the string is %v", string(abc))
	var res bool
	for i, j := 0, len(abc)-1; i < j; i, j = i+1, j-1 {
		if abc[i] != abc[j] {
			//fmt.Printf("\nthis is not palindrome")
			res = false
			break
		} else {
			res = true
		}

	}

	if res == true {
		fmt.Printf(" This is palindrome")
	} else {
		fmt.Printf(" this is not palindrome")
	}
}
