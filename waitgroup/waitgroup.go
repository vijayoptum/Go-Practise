package main

import (
	"fmt"
	"sync"
)

func waiting(wg *sync.WaitGroup) {
	fmt.Print("\nthis is function from waiting function")
	wg.Done()
}

func main() {

	fmt.Print("\nthis is function from starting main ")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go waiting(&wg)
	wg.Wait()
	fmt.Print("\nthis is function from end main ")

}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go dosomething(&wg)

// 	}

// 	wg.Wait()
// 	fmt.Println("Finish for loop")
// }

// func dosomething(wg *sync.WaitGroup) {
// 	fmt.Println("Do something")
// 	wg.Done()
// }
