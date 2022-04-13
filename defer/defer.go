package main

func add(a int, b int) int {
	return (a + b)
}

func sub(a int, b int) int {
	return (a - b)
}

func mul(a int, b int) int {
	return (a * b)
}

func main() {

	defer println(add(9, 5))
	defer println(sub(9, 5))
	println(mul(9, 5))

}
