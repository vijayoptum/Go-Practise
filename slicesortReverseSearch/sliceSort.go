package main

import (
	"fmt"
	"sort"
)

func intslicesort() {
	intSlice := []int{10, 5, 25, 351, 14, 9} // unsorted
	fmt.Println("Slice of integer BEFORE sort:", intSlice)
	sort.Ints(intSlice)
	println(sort.IntsAreSorted(intSlice))
	fmt.Println("Slice of integer AFTER  sort:", intSlice)
}

func stringslicesort() {
	istrSlice := []string{"Jamaica", "Estonia", "Indonesia", "Hong Kong"} // unsorted
	fmt.Println("Slice of integer BEFORE sort:", istrSlice)
	sort.Strings(istrSlice)
	fmt.Println("Slice of integer AFTER  sort:", istrSlice)
}

func intslicesearch() {
	intSlice := []int{10, 5, 25, 351, 14, 9} // unsorted
	abc := sort.SearchInts(intSlice, 25)
	println(abc)
}

func search() {
	intSlice := []int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5} // unsorted
	abc := sort.Search(len(intSlice), func(i int) bool { return intSlice[i] <= 52 })
	println(abc)
}

func Slice() {
	s := []int{9, 22, 54, 33, -10, 40} // unsorted
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)                                                              // sorted
	fmt.Println("Length of Slice: ", sort.IntSlice.Len(s))                      // 6
	fmt.Println("40 found in Slice at position: ", sort.IntSlice(s).Search(40)) //	4
	fmt.Println("82 found in Slice at position: ", sort.IntSlice(s).Search(82)) //	6
	fmt.Println("6 found in Slice at position: ", sort.IntSlice(s).Search(6))

}

func main() {
	//intslicesort()
	//stringslicesort()
	//intslicesearch()
	//search()
	Slice()
}
