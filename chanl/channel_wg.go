package main

import (
	"fmt"
	"sync"
)

func show(chan1 chan string, wg *sync.WaitGroup) {
	fmt.Println(len(chan1))
	for i := range chan1 {

		println(i)
	}

	wg.Done()
}

func main() {

	wg := sync.WaitGroup{}
	chn1 := make(chan string, 2)
	chn2 := make(chan string, 2)
	defer close(chn1)
	defer close(chn2)
	wg.Add(1)
	go show(chn1, &wg)
	chn1 <- "GAurav"
	chn1 <- "GAurav1"
	chn1 <- "GAurav2"
	chn1 <- "GAurav3"
	chn1 <- "GAurav4"
	chn1 <- "GAurav5"
	chn1 <- "GAurav6"

	wg.Wait()
	wg.Add(1)
	go show(chn2, &wg)
	chn2 <- "Singh"

	wg.Wait()

}
