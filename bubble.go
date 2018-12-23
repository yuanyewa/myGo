package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func BubbleSort(x []int) {
	for j := 0; j <len(x)-1; j++ {
		for i, _ := range x[0:len(x)-1-j] {
			if x[i] > x[i+1] {
				Swap(x, i)
			}
		}
	}
}
func Swap(x []int, i int) {
	x[i], x[i+1] = x[i+1], x[i]
}
func main() {
	fmt.Println("please enter a sequence of up to 10 integers, separated by ,")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	sls := make([]int, 0, 10)
	for _, s := range strings.Split(str, ",") {
		i, _ := strconv.Atoi(s)
		sls = append(sls, i)
	}
	BubbleSort(sls)
	fmt.Println("Sorted: ", sls)
}