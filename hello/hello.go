package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Printf("Hello, World!\n")
	fmt.Println(quote.Go())

	message := greetings.Hello("Gladys")
	fmt.Printf(message)
}
