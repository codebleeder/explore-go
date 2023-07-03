package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func randomFormat() string {
	// slice:
	formats := []string{
		"Hi, %v. welcome",
		"great to see you %v",
		"namaste %v",
	}

	return formats[rand.Intn(len(formats))]
}
