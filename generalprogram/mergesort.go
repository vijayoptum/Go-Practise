package main

import "fmt"

func merge(s, k []int) []int {
	merged := make([]int, 0, len(s)+len(k))

	for len(s) > 0 || len(k) > 0 {
		if len(s) == 0 {
			return append(merged, k...)
		} else if len(k) == 0 {
			return append(merged, s...)
		} else if s[0] < k[0] {
			merged = append(merged, s[0])
			s = s[1:]
		} else {

			merged = append(merged, k[0])
			k = k[1:]
		}

	}

	return merged
}

func mergehalf(xyz []int) []int {

	if len(xyz) <= 1 {
		return xyz
	}
	//done := make(chan bool)
	mid := len(xyz) / 2
	var left []int
	//go func() {
	left = mergehalf(xyz[:mid])
	//}()
	right := mergehalf(xyz[mid:])
	//<-done
	return merge(left, right)
}

func main() {

	abc := []int{56, 78, 54, 34, 1, 9, 53}
	fmt.Printf("merge sort of \n %d is \n %d", abc, mergehalf(abc))
}
