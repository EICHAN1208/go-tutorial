package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Gladys")
	// fmt.Println(message)

	// fmt.Printf("Hello, World!\n")
	// fmt.Println(quote.Go())

	// message := greetings.Hello("Gladys")
	// fmt.Printf(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
