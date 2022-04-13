package main

import (
	"fmt"
	"sort"
)

func crushingstone(abc []int) int {
	fmt.Println(abc)
	count := 0
	for i := 0; i < len(abc); i++ {
		if i < len(abc)-1 && abc[i] == abc[i+1] {
			count = count + 0
		} else {
			count = abc[i] - abc[i+1]
		}
	}
	return count
}

func main() {
	abc := []int{12, 34, 24, 53, 64, 73, 24, 64, 76, 69}
	//jkl := []string{"G", "A", "U", "R", "A", "V", "S"}

	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(cap(abc))
	fmt.Println(append(abc, 99))
	fmt.Println("after appending ", abc)
	fmt.Println(abc[len(abc)-1])
	xyz := abc
	xyz = append(xyz, 90)
	fmt.Println("after appending ", xyz)
	sort.Ints(abc)
	fmt.Println("after sorting ", abc)
	println(xyz == nil)
	fmt.Println(sort.IntSlice(abc).Search(69))
	fmt.Println(sort.SearchInts(abc, 12))
	for i, j := 0, len(abc)-1; i < j; i, j = i+1, j-1 {
		abc[i], abc[j] = abc[j], abc[i]
	}
	fmt.Println(abc)

	var zz int = crushingstone(abc)
	println(zz)

}
