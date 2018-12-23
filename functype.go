package main

import (
	"fmt"
	"strconv"
)

func add1(x int) int {
	return x + 1
}
func applyFn(f func (int) int, x int) int {
	return f(x)
}
func makefunc(x int, y int) func (int) int {
	return func(z int) int {
		return z + x + y
	}
}

func prtSlice(x []int) func () {
	return func () {
		fmt.Println(x)
	}
}
func vara(x ...int) {
	s := 0
	for _, v := range x {
		s += v
	}
	fmt.Println(s)
}

func fA() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func ffA() (ffB func (string) int){
	ffB = func(x string) int {
		v, _ := strconv.Atoi(x)
		return v
	}
	return
}

func main() {
	var c func (int) int
	c = add1 // variable as function
	fmt.Println(c(1))
	fmt.Println(applyFn(add1, 1)) // function as input
	fmt.Println(applyFn(func (x int) int {return x + 100}, 1)) // anonymous function
	f := makefunc(10,100) // function as output
	fmt.Println(f(1))
	sls := []int{1,2,3}
	val := 100
	defer fmt.Println("Deferred print: ", sls, val) // value to deferred function is fixed immediately; if value is a pointer, the content may still change
	val ++
	p := prtSlice(sls)
	p()
	sls[0] = 1000
	p() // note here, printed [0] is 1000, meanning function p has access to sls parameter through it's constructing function. p is not fixed by the time it's created.
	vara([]int{3,4,5}...) // varadic function: in definition: var ...Type; when call: slice...
	fB := fA()
	fmt.Println(fB())
	fmt.Println(fB())

	ffB := ffA()
	fmt.Println(ffB("30000"))
}