// aaabbccdaabb

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countofnumber(abc []rune) []string {
	xyz := []string{}
	count := 1
	for i := 0; i < len(abc); i++ {
		if i < len(abc)-1 && abc[i] == abc[i+1] {
			count = count + 1

		} else {
			//fmt.Println(string(abc[i]), count)
			xyz = append(xyz, string(abc[i]))
			xyz = append(xyz, strconv.Itoa(count))

			count = 1
		}

	}
	return xyz
}

func nonrepeating(abc []rune) string {
	xyz := []string{}
	count := 1
	for i := 0; i < len(abc); i++ {
		if i < len(abc)-1 && abc[i] == abc[i+1] {
			count = count + 1

		} else {
			//fmt.Println(string(abc[i]), count)
			xyz = append(xyz, string(abc[i]))
			xyz = append(xyz, strconv.Itoa(count))
			if count == 1 {
				return string(abc[i])
			}
			count = 1
		}

	}
	return "No value"
}

func numberofchar(abc []rune) {
	//xyz := []string{}
	count := 1

	kk := make(map[string]int)

	for i := 0; i < len(abc); i++ {
		_, isprent := kk[string(abc[i])]
		if isprent {
			//count = count + 1
			kk[string(abc[i])] = kk[string(abc[i])] + 1
		} else {
			kk[string(abc[i])] = count
			//count = 0
		}
	}

	for id, value := range kk {
		println(id, value)
	}
}

func substringInString(abc string) int {
	sub := "ab"
	count := strings.Count(abc, sub)
	return (count)
}

func main() {

	s := "aaabbccdaabb"
	abc := []rune(s)
	xyz := countofnumber(abc)
	l := nonrepeating(abc)
	fmt.Println(xyz)
	fmt.Println(l)

	numberofchar(abc)

	count := substringInString(s)
	println(count)
}
