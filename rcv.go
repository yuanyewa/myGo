package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Animal struct{
	food, locomotion, noise string
}

func (animal *Animal) Eat(){
	fmt.Println(animal.food)
}

func (animal *Animal) Move(){
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak(){
	fmt.Println(animal.noise)
}

func main(){
	var input, animalName, animalMethod string;
	var animal *Animal
	
	cow  := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake:= Animal{"mice", "slither", "hsss"}
	
	for{
		fmt.Print(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		fields := strings.Split(input," ")
		animalName = fields[0]
		animalMethod = fields[1]
		
		switch{
			case animalName == "cow":
				animal = &cow
			case animalName == "bird":
				animal = &bird
			case animalName == "snake":
				animal = &snake
		}
		
		switch{
			case animalMethod == "eat":
				animal.Eat()
			case animalMethod == "move":
				animal.Move()
			case animalMethod == "speak":
				animal.Speak()
		}

	}
}