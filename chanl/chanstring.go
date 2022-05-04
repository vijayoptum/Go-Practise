package main

import "fmt"

func main() {

	chan22 := make(chan string)

	go mychan(chan22)

	for {

		res, ok := <-chan22
		if ok == false {
			fmt.Println("Channel Close : ", ok)
			break
		}
		fmt.Println("Channel Open : ", res, ok)

	}
}

func mychan(ch chan string) {

	for i := 0; i < 4; i++ {

		ch <- "VijayBattula"

	}
	close(ch)

}
