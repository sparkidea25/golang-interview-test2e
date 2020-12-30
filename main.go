package main

import (
	"fmt"
	"strings"
)

const (
	A = "hello world"
	B = "Who"
)

func check(word string) {
	aContains := false
	bContains := false

	fmt.Print(word)
	word = strings.ToLower(word)

	if strings.Contains(strings.ToLower(A), word) {
		aContains = true
	}

	if strings.Contains(strings.ToLower(B), word) {
		aContains = true
	}

	if aContains && !bContains {
		fmt.Println(" exists in A and not B")
	}

	if !aContains && bContains {
		fmt.Println(" exists in B and not A")
	}

	if !aContains && !bContains {
		fmt.Println(" doesnt exist in both A and B")
	}
}

func main() {
	check("Hello")
	check("Who")
}
