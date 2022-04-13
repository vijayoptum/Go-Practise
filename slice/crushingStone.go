package main

import (
	"fmt"
)

func crushingstone(abc []int) int {
	fmt.Println(abc)
	count := abc[0]
	for i := 1; i < len(abc); i++ {
		if i < len(abc)-1 && abc[i] == count {

			count = abc[i+1]
			i = i + 1

		} else { //if i != len(abc)-1 {
			count = count - abc[i]
		} //else {
		//count = count - abc[]
		//}
	}
	return count
}
func main() {
	abc := []int{10, 10, 50, 35, 40}
	fmt.Println(abc)
	// for i, j := 0, len(abc)-1; i < j; i, j = i+1, j-1 {
	// 	abc[i], abc[j] = abc[j], abc[i]
	// }
	// fmt.Println(sort.IntSlice(abc))
	var zz int = crushingstone(abc)
	fmt.Println(zz)

}
