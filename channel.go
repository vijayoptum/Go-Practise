package main

import "fmt"

func main() {

	var testchan chan int

	fmt.Println("value of channel is :", testchan)
	fmt.Printf("Type of channel is %T :", testchan)

	testchan2 := make(chan int)

	fmt.Println("\n value of channel22 is :", testchan2)
	fmt.Printf("\n Type of channel22 is %T :", testchan2)

}
