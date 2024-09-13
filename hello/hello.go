package main

import (
	"fmt"
	"github.com/abdelrahmanMsaeed/Go-Modules/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings Module: ")
	//log.SetFlags(1)

	// Request a greeting message.
	message, err := greetings.Hello("Saeed")

	// Request a greeting messages
	//  A slice of names
	names := []string{"saeed", "Ahmed", "Ali"}
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	log.Println("message has been sent successfully")

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

}
