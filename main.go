package main

import (
	"fmt"

	"github.com/misrab/learn/api"
)


func main() {

	for {
		fmt.Print("Enter text: ")
		var input string
		fmt.Scanln(&input)
		fmt.Println(input)
	}
	
}