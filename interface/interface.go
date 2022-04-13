package main

import "fmt"

type dimension struct {
	length int
	width  int
}

type area interface {
	tarea() int
	parea() int
}

func (m dimension) tarea() int {
	return m.length * m.width

}
func (m dimension) parea() int {
	return m.length * m.width

}

func main() {
	var xyz area
	var i, j int
	fmt.Print("Please provide the value of dimesion")
	fmt.Scan(&i, &j)
	xyz = dimension{i, j}
	fmt.Printf("the value of area is %d", xyz.tarea())
}
