package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	message, err := greetings.Hello("Sharad")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
