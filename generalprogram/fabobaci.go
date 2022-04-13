package main

import "fmt"

func fabonaci(s int) {
	//z := 0
	x := 0
	y := 1
	//k := x + y
	for i := 0; i < s; i++ {
		if i == 0 {
			fmt.Printf("\n %d", x)
		} else if i == 1 {
			fmt.Printf("\n %d", y)
		} else {
			k := x + y
			x = y
			y = k
			fmt.Printf("\n %d", k)
		}

	}

}

func main() {
	var i int
	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&i)
	//i := 10
	fabonaci(i)

}

//0,1,1,2,3,5,8,13
