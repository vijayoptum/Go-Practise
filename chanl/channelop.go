// package main

// import (
// 	"fmt"
// 	"time"
// )

// func myfunc(ch chan int, ch1 chan int) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(257 + <-ch)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(267 + <-ch1)
// }
// func main() {
// 	fmt.Println("start Main method")
// 	// Creating a channel
// 	var ch1 chan int
// 	ch := make(chan int)

// 	go myfunc(ch, ch1)
// 	ch1 <- 56
// 	ch <- 23

// 	fmt.Println("End Main method")
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
	//wg.Wait()
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	//var ch chan int
	ch := make(chan int, 3)
	wg := sync.WaitGroup{}

	for i := 0; i < cap(ch); i++ {
		//fmt.Println(ch)
		go myfunc(ch, &wg)
		wg.Add(1)
		ch <- i + 32

		//fmt.Println(ch)
	}
	//close(ch)
	//wg.Wait()
	println(len(ch))

	// for res := range ch {
	// 	wg.Add(1)
	// 	fmt.Printf("\nthe value of res%d\n", (res))
	// 	wg.Done()

	// }
	//
	wg.Wait()
	close(ch)
	fmt.Println("End Main method")

}
