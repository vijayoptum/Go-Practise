package main

import (
	"fmt"
	"time"
)

func display(s string) {

	for i := 0; i < 3; i++ {

		fmt.Print("\n", s)
		time.Sleep(1000 * time.Microsecond)
	}

}

func main() {

	go display("welcome")
	time.Sleep(1000 * time.Microsecond)
	display("\nGaurav Singh")

}
