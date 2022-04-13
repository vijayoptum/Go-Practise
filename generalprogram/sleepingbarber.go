package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_BARBERS  = 2
	LOBBY_SIZE   = 3
	HAIRCUT_TIME = 20
)

func main() {
	customers := make(chan *Customer)
	go CustomerProducer(customers)
	go BarberShop(customers)
	time.Sleep(2 * time.Second)
}

type Barber struct {
	Id int
}

func (b Barber) String() string {
	return fmt.Sprintf("%d", b.Id)
}

type Customer struct {
	Id int
}

func (c Customer) String() string {
	return fmt.Sprintf("%d", c.Id)
}

func CustomerProducer(customers chan *Customer) {
	customerCounter := 1
	for {
		rnd := rand.Intn(10)
		time.Sleep(time.Duration(rnd) * time.Millisecond)
		customers <- &Customer{customerCounter}
		customerCounter += 1
	}
}

func CutHair(barber *Barber, customer *Customer, finished chan *Barber) {
	time.Sleep(HAIRCUT_TIME * time.Millisecond)
	fmt.Printf("Barber %s done with customer %s.\n", barber, customer)
	finished <- barber
}

func BarberShop(customers <-chan *Customer) {
	barbers := []*Barber{}
	lobby := []*Customer{}
	syncBarberChan := make(chan *Barber)

	for i := 0; i < NUM_BARBERS; i++ {
		barbers = append(barbers, &Barber{i + 1})
	}

	for {
		select {
		// take a customer off of the customer channel
		case customer := <-customers:
			// if there are no barbers, either sit in the lobby or leave
			if len(barbers) == 0 {
				if len(lobby) < LOBBY_SIZE {
					lobby = append(lobby, customer)
					fmt.Printf("Customer %s seated in lobby.\n", customer)
				} else {
					fmt.Printf("Lobby full, customer %s left.\n", customer)
				}
			} else {
				barber := barbers[0]
				barbers = barbers[1:]
				fmt.Printf("Customer %s goes to barber %s.\n", customer, barber)
				go CutHair(barber, customer, syncBarberChan)
			}
			fmt.Printf("Lobby: %+v\n", lobby)
		// take a barber off of the sync barber channel
		case barber := <-syncBarberChan:
			// if there are customers in the lobby, grab one, otherwise take a nap
			if len(lobby) > 0 {
				customer := lobby[0]
				lobby = lobby[1:]
				fmt.Printf("Customer %s goes to barber %s.\n", customer, barber)
				go CutHair(barber, customer, syncBarberChan)
			} else {
				barbers = append(barbers, barber)
				fmt.Printf("Barber %s idle.\n", barber)
			}
			fmt.Printf("Lobby: %+v\n", lobby)
		}
	}
}
