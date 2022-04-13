package main

import "fmt"

func datatype(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Printf("\nthe value is of integer type %d", a.(int))
		println(a.(int))
	case string:
		fmt.Printf("\nthe value is of string type %s", a.(string))
	case float64:
		fmt.Printf("\nthe value is of float type	%T", a.(float64))
	case bool:
		fmt.Printf("\nthe value is of bool type	%T", a.(bool))
	default:
		fmt.Print("\nthe value is of other type")
	}

}

func value1(m interface{}) {

	if m == 10 {
		print("got the value")
	} else {
		print("not got the value")
	}

}

func myfun(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}

func main() {
	// datatype("Gaurav")
	// datatype(123)
	// datatype(true)
	// datatype(12.4500)
	//value1(20)

	var a1 interface{} = 98.76

	myfun(a1)
}
