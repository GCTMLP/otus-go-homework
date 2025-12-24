package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func ReverseString(incomingString string) string {
	return reverse.String(incomingString)
}

func main() {
	// Place your code here.
	HelloString := "Hello, OTUS!"
	ReversedString := ReverseString(HelloString)
	fmt.Println(ReversedString)
}
