package main

import (
	"fmt"
	"sort"
)

func main() {

	abc := [10]int{12, 34, 24, 53, 64, 73, 24, 64, 76}
	xyz := abc
	abc[5] = 10
	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(xyz)
	fmt.Println(xyz == abc)
	sort.Ints(abc[:])
	fmt.Println(abc)

	for _, j := range abc {
		println(j)
	}

	hh := [5]int{}
	fmt.Println(hh)
}
