package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// fmt.Printf("Hello, World!\n")
	// fmt.Println(quote.Go())

	// message := greetings.Hello("Gladys")
	// fmt.Printf(message)
}
