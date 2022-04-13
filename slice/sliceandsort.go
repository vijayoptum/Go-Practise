package main

import (
	"fmt"
	"sort"
)

func main() {

	slice1 := []int{8, -5, -6, 4, 1, 9, 3, 3}
	slice2 := []string{"gaurav", "gau", "bavita", "aradhya"}

	fmt.Printf("\ninterger slice before sort %d", slice1)
	sort.Ints(slice1)
	fmt.Printf("\ninterger slice after sort %d", slice1)
	fmt.Printf("\nstring slice before sort %s", slice2)
	sort.Strings(slice2)
	fmt.Printf("\nstring slice after sort %s", slice2)
	if sort.IntsAreSorted(slice1) == true {
		fmt.Printf("it is sorted")
	}

}
