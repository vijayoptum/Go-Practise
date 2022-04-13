package main

func reciever(c <-chan int) {
	for v := range c {
		println(v)
	}

}

func sender() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {

			c <- i

		}
		close(c)
	}()

	return c
}

func main() {
	c := sender()
	reciever(c)

}
