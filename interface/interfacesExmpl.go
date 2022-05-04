package main

import "fmt"

type Bird interface {
	Fly()
	Fly1()
}

type Eagle struct {
	name string
}

func (e Eagle) Fly1() {

	fmt.Println("Eagle is flying over the cloud FLY1")
}

func (e Eagle) Fly() {

	fmt.Println("Eagle is flying over the cloud")
}

type Pigeon struct {
}

func (p Pigeon) Fly() {
	fmt.Println("Pigeon is flying on normal height")
}

func (p Pigeon) Fly1() {
	fmt.Println("Pigeon is flying on normal height")
}

type Penguin struct {
}

func (p Penguin) Fly() {
	fmt.Println("Penguin cannot fly")
}

func (p Penguin) Fly1() {
	fmt.Println("Penguin cannot fly")
}
func flyNow(b Bird) {
	b.Fly()
	b.Fly1()

	// eg.Fly()
	// pg.Fly()
	// pe.Fly()
}

// flynow aney functipon elanti type(struct) ni ina input ga thesukuntadhi

func main() {

	// eg := Eagle{"mybirdeagle"} // eg type  is implementing a struct  eg is object/value
	// pg := Pigeon{}
	// pe := Penguin{}

	// eg.Fly() // eg struct value ni use chesi fly anna method in implelment cheyagalu
	// pg.Fly()
	// pe.Fly()

	// eg.Fly()
	// eg.Fly1()

	flyNow(Eagle{})
	flyNow(Pigeon{})
	flyNow(Penguin{})

}
