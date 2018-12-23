package main

import "fmt"

func main() {
	switch i := 3; i {
	case 3, 10:
		fmt.Println("This is 3")
	case 4:
		fmt.Println("This is 3")
	case 5:
		fmt.Println("This is 3")
	}
	switch {
	case true:
		fmt.Println("this is rue")
	case false:
		fmt.Println("this is not true")
	}
}
