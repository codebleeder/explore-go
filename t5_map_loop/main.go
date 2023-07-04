package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	names := []string{"sharad", "shweta", "surabhi"}
	res, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
