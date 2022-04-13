package main

import "fmt"

func reverse(s int) int {
	//l := 2
	res := 0
	for s > 0 {
		reminder := s % 10
		res = (res * 10) + reminder
		s /= 10

	}
	return (res)
}

func reversestring(s string) string {
	//l := 2
	abc := []byte(s)
	for i, j := 0, len(abc)-1; i < j; i, j = i+1, j-1 {
		abc[i], abc[j] = abc[j], abc[i]

	}
	return (string(abc))
}

func main() {
	var i int
	fmt.Print("\nEnter the number of terms : ")
	fmt.Scan(&i)
	//i := 10

	fmt.Print("\nreverse is ", reverse(i))

	var j string
	fmt.Print("\nEnter the string : ")
	fmt.Scan(&j)
	fmt.Print("\nreverse is ", reversestring(j))

}

//1234 === 4321
