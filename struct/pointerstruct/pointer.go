package main

import "fmt"

// defining a structure
type Employee struct {
	firstName, lastName string
	age, salary         int
}

type hr struct {
	details Employee
}

func main() {

	// passing the address of struct variable
	// emp8 is a pointer to the Employee struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

	// (*emp8).firstName is the syntax to access
	// the firstName field of the emp8 struct
	fmt.Println("First Name:", (emp8).firstName)
	fmt.Println("Age:", (emp8).age)
}
