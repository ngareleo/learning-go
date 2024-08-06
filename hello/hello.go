package main

import (
	"fmt"

	"rsc.io/quote"

	"learninggo/greetings"
)

func HelloWorld () {
	fmt.Println(quote.Go())
}

func main() {
	HelloWorld()
	message := greetings.Hello("Leo")
	fmt.Println(message)
}