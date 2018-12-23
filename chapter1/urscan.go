package main

import (
	"fmt"
)

func GenDisplaceFn2(a, v, s float64) func(t float64) float64{
	return func(t float64) float64{
		return 0.5*a*t*t + v*t + s
	}
}

func main(){
	var s, t, a, v float64
	
	fmt.Print("Value of acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Value of velocity: ")
	fmt.Scan(&v)
	fmt.Print("Value of initial displacement: ")
	fmt.Scan(&s)
	fmt.Print("Value of time: ")
	fmt.Scan(&t)
	
	fn := GenDisplaceFn2(a, v, s)
	fmt.Println(fn(t))
}