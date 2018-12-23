package main

import (
	"fmt"
	"time"
)

var ary = []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5, 13, 14, 11, 13, -1}
var target = 10

func main() {
	start := time.Now()
	fmt.Println(ary)
	for l := 0; l < 2; l++ {
		sumtwo()
	}
	fmt.Println(time.Since(start).Seconds())
}

func pair_sum() {
	seen := make(map[int]bool)
	output := make(map[int]int)
	for _, num := range ary {
		t := target - num
		if _, ok := seen[t]; !ok {
			seen[t] = true
		} else {
			if num > t {
				num, t = t, num
			}
			output[num] = t
		}
	}
	// fmt.Println(output)
}
func sumtwo() {
	pairs := make(map[int]int)
	for i, a := range ary {
		for _, b := range append(ary[0:i], ary[i+1:len(ary)]...) {
			if a+b == target {
				if a > b {
					a, b = b, a
				}
				pairs[a] = b
				break
			}
		}
	}
	fmt.Println((pairs))
}
