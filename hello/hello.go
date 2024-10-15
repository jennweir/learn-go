package main

import (
	"fmt"
	"log"

	// For production use, you’d publish the example.com/greetings module from its repository
	// (with a module path that reflected its published location), where Go tools could find it to download it
	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// message1, err1 := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }

	// If no error was returned, print the returned message
	// to the console.

	fmt.Println("Hello world!")
	fmt.Println(quote.Go())
	message2, err2 := greetings.Hello("Gladys")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)
}
