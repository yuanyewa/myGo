package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
)

type name struct {
	Fname string
	Lname string
}
func main() {
	nameStruct := make([]name, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input filename: ")
	filename,_ := reader.ReadString('\n')
	filename = strings.TrimSuffix(filename, "\r\n")
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		newName := name{Fname:fields[0], Lname: fields[1]}
		nameStruct = append(nameStruct, newName)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, name := range nameStruct {
		fmt.Print("Fname: ", name.Fname, "; Lname: ", name.Lname, "\n")
	}
}