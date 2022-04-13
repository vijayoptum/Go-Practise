package main

import "fmt"

func main() {
	var arr [5]int

	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4

	for i := range arr {
		println(arr[i])
	}
	fmt.Printf("\nonly arr%d", arr)
	fmt.Printf("\nonly single%d", arr[0])

	arr1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("\nonly arr1%d", arr1)
	fmt.Println("\nLength of the array  is:", len(arr1))

	//arr2 := [5]int{0, 1, 2, 3, 4}

	arr3 := &arr
	arr[4] = 1000

	if *arr3 == arr {
		fmt.Printf("both are equal")
		fmt.Printf("array 3 data %d", *arr3)

	} else {
		fmt.Printf("both are unequal")
	}

}
