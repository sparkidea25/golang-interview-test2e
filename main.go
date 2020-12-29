package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = "welcome to my world of Golang"
	var b = "welcome to the world of Fintech and Blockchain"

	input := "world"
	inpute := "Fintech"

	if strings.Contains(a, input) {
		fmt.Println("welcome a")
	}
	if strings.Contains(b, input) {
		fmt.Println("Welcome b")
	}

	if strings.Contains(a, inpute) {
		fmt.Println("Welcome")
	} else {
		fmt.Println("No it doesnt")
	}

	if strings.Contains(b, inpute) {
		fmt.Println("welcome")
	} else {
		fmt.Println("no it doesnt")
	}
}
