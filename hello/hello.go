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
	// Had to run go mod edit -replace learninggo/greetings=../greetings 
	// to point go to the greetings module
	// then => go mod tidy to resolve dependencies or build the tree (not sure yet)
	HelloWorld()
	message := greetings.Hello("Leo")
	fmt.Println(message)
}