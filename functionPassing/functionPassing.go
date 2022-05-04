package main

func add(abc [5]int, size int) int {
	var sum int

	for i := range abc {
		sum += abc[i]

	}
	return sum

}

func main() {

	arr2 := [...]int{10, 1, 1, 3, 4}
	length := len(arr2)

	c := add(arr2, length)
	println(c)

}
