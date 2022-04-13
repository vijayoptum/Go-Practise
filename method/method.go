package main

import "fmt"

type data int

type family struct {
	name   string
	age    int
	salary float32
}

func (abc family) desc() {
	fmt.Printf("\nname is : %s", abc.name)
	fmt.Printf("\nage is : %d", abc.age)
	fmt.Printf("\nsalary  is : %f", abc.salary)

}

func (a data) mul(b data) data {

	return (a * b)
}

func main() {

	//fmt.Printf("this is my first program")

	var xyz family
	xyz.name = "gaurav"
	xyz.age = 31
	xyz.salary = 14.50

	xyz.desc()

	value1 := data(10)
	value2 := data(20)

	fmt.Printf("\nthe mul value is : %d", value1.mul(value2))
	//println(value1.mul(value2))

}
