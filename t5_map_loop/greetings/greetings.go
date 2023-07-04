package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hellos(names []string) (map[string]string, error) {
	// instantiate map:
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

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
