package main

import (
	"fmt"
	"log"
	"t3_error_handling"
)

func main() {
	callHello("sharad")
	callHello("")
}

func callHello(name string) {
	message, err := t3_error_handling.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
