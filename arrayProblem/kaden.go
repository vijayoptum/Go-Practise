package main

import (
	"fmt"
	"math"
)

func main() {
	abc := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println(abc)
	xyz := []int{}

	max_so_far := math.MinInt16
	max_end := 0

	for i := 0; i < len(abc); i++ {
		max_end = max_end + abc[i]
		xyz = append(xyz, abc[i])

		if max_so_far < max_end {
			max_so_far = max_end

		}

		if max_end < 0 {
			max_end = 0
			xyz = []int{}
		}

	}

	println(max_so_far)
	fmt.Println(xyz)

}
