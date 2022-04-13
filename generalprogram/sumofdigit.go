package main

import "fmt"

func sumOfDigit(s int) int {
	//l := 2
	res := 0
	for s > 0 {
		reminder := s % 10
		res = res + reminder
		s /= 10
		println(reminder)
		println(res)

	}
	return (res)
}

func sumOfDigitByNine(i int) int {
	if i == 0 {
		return 0
	}

	if i%9 == 0 {
		return 9
	} else {
		return (i % 9)
	}

}

func main() {
	var i int
	fmt.Print("\nEnter the number of terms : ")
	fmt.Scan(&i)
	//i := 10

	fmt.Print("\n sum of digit is ", sumOfDigit(i))
	fmt.Print("\n sum of digit is ", sumOfDigitByNine(i))

}

//1234 === 4321
