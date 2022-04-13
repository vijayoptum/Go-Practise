package main

import (
	"fmt"
)

func maxdif(abc []int) int {
	//max_ele := abc[1]
	min_ele := abc[0]
	max_dif := abc[1] - abc[0]

	for i := 1; i < len(abc); i++ {
		if abc[i]-min_ele > max_dif {
			max_dif = abc[i] - min_ele
		}
		if abc[i] < min_ele {
			min_ele = abc[i]
		}
	}
	return max_dif
}

func main() {

	abc := []int{1, 2, 90, 10, 110}
	fmt.Println(abc)

	sum := maxdif(abc) // when we have to find only 1 combination
	println(sum)
}
