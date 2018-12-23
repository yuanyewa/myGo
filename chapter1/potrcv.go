package main

import (
	"fmt"
)
type myT struct {
	x int
}

func (m *myT) Double() {
	m.x *= 2
}
func (m *myT) Double2() {
	(*m).x *= 2
}
func (m myT) Double3() {
	m.x *= 2
}

func main() {
	m := myT {5}
	m.Double3()
	fmt.Println(m.x)
	m.Double()
	fmt.Println(m.x)
	m.Double2()
	fmt.Println(m.x)
	(&m).Double()
	fmt.Println(m.x)
	(&m).Double2()
	fmt.Println(m.x)
}
