package main

import "fmt"

func fabonaci(num int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		ch <- x
		x, y = y, x+y
		//println(<-ch)
	}
	close(ch)
}

func main() {
	var n int
	fmt.Printf("Please enter the number for which fabonaci series need to be printed")
	fmt.Scan(&n)
	ch := make(chan int, n)
	fmt.Printf("Starting of main and fabonaci series\n")
	go fabonaci(cap(ch), ch)
	for x := range ch {
		fmt.Println(x)
	}
	fmt.Printf("fabonaci series is completed")

}
