package main

import (
	"fmt"
	"log"

	"github.com/AlexDJensen/learning-go/pkg/greetings"
)

func main() {
	//Setup the logger below:
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Get a greeting if possible
	//message, err := greetings.Hello("Alexander")
	//Use time to determine what input should be used:
	var input string
	// if time.Now().Unix()%10 < 2 {
	// 	input = ""
	// } else {
	input = "Alexander"
	// }

	message, err := greetings.Hello(input)

	//Exit if the error is not nil
	if err != nil {
		log.Fatal(err)
	}

	//If no error, print message
	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Println("Greeting for :", name, "is as follows: ", message)
	}
}
