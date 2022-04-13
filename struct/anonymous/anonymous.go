package main

import "fmt"

type family struct {
	name string
	age  int
	float32
}

func main() {

	//fmt.Printf("this is my first program")

	xyz := family{"gaurav", 31, 14.50}

	// fmt.Printf("details before change")
	// fmt.Printf("\nname before change %s", xyz.name)
	// fmt.Printf("\nname before change %d", xyz.age)
	// fmt.Printf("\nname before change %f", xyz.salary)
	fmt.Println(xyz)

}
