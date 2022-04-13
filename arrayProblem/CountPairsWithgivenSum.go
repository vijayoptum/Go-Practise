package main

import (
	"fmt"
	"sort"
)

func usingtwopointer(abc []int, sum int) int {
	sort.Ints(abc)
	fmt.Println(sort.IntSlice(abc))
	count := 0
	i := 0
	j := len(abc) - 1
	for i < j {
		if (abc[i] + abc[j]) == sum {
			count = count + 1
			println(i, j)
			//return i, j

		}
		if (abc[i] + abc[j]) < sum {
			i++
		} else {
			j--
		}

	}
	return count //i, j
}

// for (abc[i]+abc[j]) != sum &&  {
// 	if (abc[i] + abc[j]) < sum {
// 		i++
// 	} else {
// 		j--
// 	}
// }

func main() {

	abc := []int{4, 7, 8, 3, 1}
	fmt.Println(abc)

	sum := 11
	println(sum)
	count := usingtwopointer(abc, sum) // when we have to find only 1 combination
	println(count)
}
