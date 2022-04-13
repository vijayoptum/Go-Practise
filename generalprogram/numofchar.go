package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "aaarrfllasarrlffff"

	//single character search

	rese1 := strings.Count(s, "a")
	fmt.Printf("\nsingle character a in string%d", rese1)

	// res := strings.Count(s, "a")
	// println(res)
	abc := []rune(s)
	// ghi := []byte(s)
	//var xyz []rune

	fmt.Printf("the rune value %d", (abc))
	// fmt.Printf("the Byte value %d", (ghi))
	// sort.Slice(abc, func(i, j int) bool {
	// 	return abc[i] < abc[j]
	// })

	fmt.Printf("\nthe rune value %d", (abc))
	count := 1
	xyz := make(map[string]int)
	commu := make([]string, 0, 10)
	for i := 0; i < len(abc); i++ {
		if i < len(abc)-1 && abc[i] == abc[i+1] {

			count = count + 1
		} else {
			//xyz = (abc[i])
			fmt.Printf("\nthe value %s and count is %d", string(abc[i]), count)
			//r := string(abc[i])
			s := strconv.Itoa(count)
			commu = append(commu, s)
			commu = append(commu, string(abc[i]))

			count = 1
		}

	}
	fmt.Println("\nthe value of map is ", xyz)
	fmt.Println("\nthe value of slice is ", strings.Join(commu, ""))
	// for i, j := range xyz {
	// 	println(i, j)
	// }

}
