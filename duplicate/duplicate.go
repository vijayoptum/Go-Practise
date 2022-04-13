package main

import "fmt"

func duplicateInArray(arr []int) []int {
	visited := make(map[int]bool, 0)
	var abcd []int
	for i := 0; i < len(arr); i++ {

		if visited[arr[i]] == true {
			abcd = append(abcd, arr[i])
			return abcd // remove this if all duplicate is needed
		} else {
			visited[arr[i]] = true
		}
	}
	return nil // return nil if all values are needed
}

func main() {
	//var abc []int
	InArray := []int{1, 4, 7, 2, 2, 4, 5, 4, 5, 1}
	fmt.Printf("Array is %v", InArray)
	abcd := duplicateInArray(InArray)
	//abc = append(abc, a)
	fmt.Println(abcd)
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3}))
}
