// package main

// import "fmt"

// // Main function
// func main() {

// 	// Creating a channel
// 	// Using make() function
// 	mychnl := make(chan string, 10)

// 	// Anonymous goroutine
// 	go func() {
// 		mychnl <- "GFG"
// 		mychnl <- "gfg"
// 		mychnl <- "Geeks"
// 		mychnl <- "GeeksforGeeks"
// 		close(mychnl)
// 	}()
// 	fmt.Println(cap(mychnl))
// 	fmt.Println(len(mychnl))
// 	// Using for loop
// 	for res := range mychnl {
// 		fmt.Println(res)
// 	}
// }

package main

import (
	"fmt"
	"sync"
)

func myfunc(ch chan int, wg *sync.WaitGroup) {
	x := <-ch
	fmt.Println(257 + x)
	wg.Done()
	//close(ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	//var ch chan int
	ch := make(chan int, 3)
	wg := sync.WaitGroup{}

	go func() {
		go myfunc(ch, &wg)
		ch <- 23
		ch <- 33
		ch <- 43
		close(ch)
	}()
	println(cap(ch))
	println(len(ch))
	for res := range ch {
		println(res)
	}

	// for i := 0; i < cap(ch); i++ {
	// 	//fmt.Println(ch)
	// 	go myfunc(ch, &wg)
	// 	wg.Add(1)
	// 	ch <- i

	// 	//fmt.Println(ch)
	// }
	//close(ch)

	wg.Wait()
	//close(ch)
	fmt.Println("End Main method")

}
