package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger
	// 1. log entry prefix
	// 2. flag to disable timestamp, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Saul", "Kim", "Mike", "Howard"}

	// request greeting messages for the names
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// if no error, print the messages
	fmt.Println(message)
}