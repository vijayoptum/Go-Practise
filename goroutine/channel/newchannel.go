package main

import (
	"sync"
)

func show(chan1 chan string, wg *sync.WaitGroup) {
	for i := range chan1 {

		println(i)

	}

	wg.Done()
}

func main() {
	i := 0
	wg := sync.WaitGroup{}
	chn1 := make(chan string)
	chn2 := make(chan string, 2)

	wg.Add(1)
	go show(chn1, &wg)

	for i < 5 {
		chn1 <- "gaurav"
		i++

	}

	close(chn1)
	wg.Wait()

	wg.Add(1)
	go show(chn2, &wg)
	chn2 <- "Singh"
	close(chn2)
	wg.Wait()

}
