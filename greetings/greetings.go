package greetings

import "fmt"

func Hello(name string) string {
	// Return a greeting that embeds the name in message.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
