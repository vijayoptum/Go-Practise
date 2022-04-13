package main

import "fmt"

func perm(abc []rune, f func([]rune)) {
	permutation(abc, f, 0)
}

func permutation(xyz []rune, h func([]rune), i int) {

	if i > len(xyz) {
		h(xyz)
		return
	}
	permutation(xyz, h, i+1)
	for j := i + 1; j < len(xyz); j++ {
		//fmt.Printf("\nfrom the 1st permutation%d%d", i, j)
		xyz[i], xyz[j] = xyz[j], xyz[i]
		//permutation(xyz, h, i+1)
		fmt.Printf("\nfrom the 2nd permutation%d%d", i, j)
		xyz[i], xyz[j] = xyz[j], xyz[i]
	}

}

func main() {

	s := "abc"
	kkk := []rune(s)

	perm(kkk, func(a []rune) { println(string(a)) })
	//perm([]rune("abc"), func(a []rune) { println(string(a)) })

}
