// // approach 1

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	pingChan := make(chan string)
// 	pongChan := make(chan string)

// 	go printer(pongChan)
// 	go pinger(pingChan)
// 	go ponger(pingChan, pongChan)

// 	// Need to block until all complete
// 	var input string
// 	fmt.Scanln(&input)
// }

// func pinger(pingChan chan<- string) {
// 	// Send a ping
// 	fmt.Println("pinger sending \"ping\"")
// 	pingChan <- "ping"
// }

// func ponger(pingChan <-chan string, pongChan chan<- string) {
// 	// receive the ping
// 	fmt.Println("ponger received", <-pingChan)

// 	// respond with a pong
// 	fmt.Println("ponger replying with \"pong\"")
// 	pongChan <- "pong"
// }

// func printer(pongChan <-chan string) {
// 	// read the pong
// 	fmt.Println("printer received", <-pongChan)
// }

//////////////////////////////////////////////////////////////

//approach 2 :
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}
func ponger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "pong"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	c := make(chan string)
	defer close(c)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

///////////////////////////////////////////////////////
// approach 3

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ping := make(chan struct{}, 1)
// 	pong := make(chan struct{}, 1)

// 	ping <- struct{}{}

// 	go play(ping, pong)

// 	time.Sleep(time.Millisecond)
// }

// func play(ping, pong chan struct{}) {
// 	for {
// 		select {
// 		case <-ping:
// 			fmt.Println("ping")
// 			pong <- struct{}{}
// 		case <-pong:
// 			fmt.Println("    pong")
// 			ping <- struct{}{}
// 		}
// 	}
// }
/////////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func pinger(c chan string, wg *sync.WaitGroup) {

// 	for i := 0; i < 5; i++ {
// 		wg.Done()
// 		c <- "ping"

// 	}
// }
// func ponger(c chan string, wg *sync.WaitGroup) {

// 	for i := 0; i < 5; i++ {
// 		wg.Done()
// 		c <- "pong"
// 		//wg.Done()
// 	}
// }
// func printer(c chan string, wg *sync.WaitGroup) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)

// 		//time.Sleep(time.Second * 1)
// 	}
// }
// func main() {
// 	var c chan string = make(chan string)
// 	defer close(c)
// 	wg := sync.WaitGroup{}
// 	wg.Add(10)
// 	go pinger(c, &wg)
// 	go ponger(c, &wg)
// 	go printer(c, &wg)
// 	wg.Wait()
// 	//c <- "ping"
// 	//time.Sleep(time.Second * 10)
// 	// var input string
// 	// fmt.Scanln(&input)
// }
