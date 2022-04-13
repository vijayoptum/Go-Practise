package main

import (
	"fmt"
	"sync"
)

const (
	numOfPhilosophers = 5
	numOfMeals        = 3
	maxEaters         = 2
)

func main() {
	chopsticks := make([]sync.Mutex, 5)
	permissionChannel := make(chan bool)
	finishEating := make(chan bool)
	go permissionFromHost(permissionChannel, finishEating)
	var wg sync.WaitGroup
	wg.Add(numOfPhilosophers)
	for i := 1; i <= numOfPhilosophers; i++ {
		go eat(i, &chopsticks[i-1], &chopsticks[i%numOfPhilosophers], &wg, permissionChannel, finishEating)
	}
	wg.Wait()
}

func eat(philosopherId int, left *sync.Mutex, right *sync.Mutex, wg *sync.WaitGroup, permissionChannel <-chan bool, finishEatingChannel chan<- bool) {
	defer wg.Done()
	for i := 1; i <= numOfMeals; i++ {
		//lock chopsticks in random order
		if RandBool() {
			left.Lock()
			right.Lock()
		} else {
			right.Lock()
			left.Lock()
		}

		fmt.Printf("waiting for permission from host %d\n", philosopherId)
		<-permissionChannel

		fmt.Printf("starting to eat %d (time %d)\n", philosopherId, i)
		fmt.Printf("finish to eat %d (time %d)\n", philosopherId, i)
		//release chopsticks
		left.Unlock()
		right.Unlock()

		//let host know I am done eating
		finishEatingChannel <- true
	}
}

func permissionFromHost(permissionChannel chan<- bool, finishEating <-chan bool) {
	ctr := 0
	for {
		if ctr < maxEaters {
			select {
			case <-finishEating:
				ctr--
			case permissionChannel <- true:
				ctr++
			}
		} else {
			<-finishEating
			ctr--
		}
	}
}

func RandBool() bool {
	//rand.Seed(time.Now().UnixNano())
	return false //rand.Intn(2) == 1
}
