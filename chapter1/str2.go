package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input file name: ")
	fileName,_ := reader.ReadString('\n')
	fileName = strings.TrimRight(fileName, "\r\n")
	
	dat, err := ioutil.ReadFile(fileName)
	s := string(make([]rune, 20, 20))
	fmt.Println(len(s))
	switch{
		case err == nil:
			fmt.Println(string(dat))
		case err != nil:
			fmt.Println("File cannot be read")
	}
}