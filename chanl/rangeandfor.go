package main

import (
	"fmt"
	"sync"
)

func myfunc(ch chan int, wg *sync.WaitGroup) {
	//x := <-ch
	fmt.Println(257 + <-ch)
	//time.Sleep(1 * time.Second)
	wg.Done()
	//close(ch)

}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	//var ch chan int
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}
	for i := 0; i < cap(ch); i++ {
		//fmt.Println(ch)
		wg.Add(2)
		go myfunc(ch, &wg)
		ch <- (i + 32)
		//time.Sleep(500 * time.Nanosecond)

		//fmt.Println(ch)
	}
	close(ch)
	//time.Sleep(2 * time.Second)
	println(len(ch))

	for res := range ch {
		//println(len(ch))
		fmt.Println(res)
		wg.Done()
		//time.Sleep(1000 * time.Nanosecond)
	}
	wg.Wait()

	fmt.Println("End Main method")

}
