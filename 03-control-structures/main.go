package main

import (
	"fmt"
	"math/rand"
)

func main() {
	///////////////////////////////////////
	///////// If statements //////////////
	///////////////////////////////////////

	// Go uses if statements like any other programming language
	age := 10
	if age < 18 {
		fmt.Println("We don't let kids in")
	} else if age > 70 {
		fmt.Println("You need atleast an instagram account to enter")
	} else {
		fmt.Println("Welcome to the clubbb!!")
	}

	// However go has something else special. Its allow you to create variables exclusive to the 
	// if block like so:
	if n := rand.Int(); n == 0 {
		fmt.Println("I can't believe we got a 0")
	} else if n > 10 {
		fmt.Println("That's more like it")
	} else if n > 100 {
		fmt.Println("I think that's a little too much")
	}

	// if we try to access n here
	// fmt.Println("Do we still have something in n? ", n) Go compiler will panic here
}