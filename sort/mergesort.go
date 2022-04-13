package main

import "fmt"

func mergesort(abc []int) []int {
	if len(abc) == 1 {
		return abc
	}
	mid := (len(abc) / 2)

	left := mergesort(abc[:mid])
	right := mergesort(abc[mid:])

	return (merge(left, right))
}

func merge(left []int, right []int) []int {
	//merged := make([]int, len(left)+len(right))
	merged := []int{}

	for len(left) > 0 || len(right) > 0 {

		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]

		}

	}
	fmt.Println(merged)
	return merged
}

func main() {

	abc := []int{7, 4, 8, 3, 54, 79, 23, 10}
	fmt.Println(abc)

	fmt.Println(mergesort(abc))

}
