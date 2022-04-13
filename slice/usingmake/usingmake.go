//func make([]T, len, cap) []T

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var slice = make([]int, 5, 8)

	//slice := []int{1, 2, 3, 4, 5}
	slice1 := slice

	for i, j := range slice {
		println(i, slice[j])
	}
	fmt.Printf("\nthe capacity is %d", cap(slice))
	fmt.Printf("\nthe length is %d", len(slice))
	fmt.Println(reflect.DeepEqual(slice, slice1))

}
