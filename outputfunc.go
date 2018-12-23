package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	return func (t float64) float64 {
		return 1.0/2 * a * t * t + v0 * t + s0
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a,v0,s0 separated by ,")
	str, _ := reader.ReadString('\n')
	sls := strings.Split(strings.TrimSuffix(str, "\r\n"), ",")
	a, _ := strconv.ParseFloat(sls[0], 64)
	v0, _ := strconv.ParseFloat(sls[1], 64)
	s0, _ := strconv.ParseFloat(sls[2], 64)
	var t float64
	fmt.Println("Please enter t:")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}