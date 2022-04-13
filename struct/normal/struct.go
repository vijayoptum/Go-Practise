package main

import "fmt"

type data int

type family struct {
	name   string
	age    int
	salary float32
}

func main() {

	//fmt.Printf("this is my first program")

	var xyz family
	xyz.name = "gaurav"
	xyz.age = 31
	xyz.salary = 14.50

	fmt.Printf("details before change")
	fmt.Printf("\nname before change %s", xyz.name)
	fmt.Printf("\nname before change %d", xyz.age)
	fmt.Printf("\nname before change %f", xyz.salary)

	abc := family{name: "bavita"}
	println(abc.name)
	println(abc.age)
	println(abc.salary)

}
