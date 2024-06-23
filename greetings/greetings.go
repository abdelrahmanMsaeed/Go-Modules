package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Random Format returns one of a set of greetings messages.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great To See You, %v!",
		"Hi, %v. Welcome!",
	}

	return formats[rand.Intn(len(formats))]
}
