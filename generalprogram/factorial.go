package main

import "fmt"

func Factorial(s int) int {
	//l := 2
	if s == 1 || s == 0 {
		return s
	} else {
		return s * Factorial(s-1)
	}

}

func main() {
	var i int
	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&i)
	//i := 10
	fmt.Print("factoial is ", Factorial(i))

}

//1*2*3
