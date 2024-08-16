package main

import (
	"errors"
	"fmt"
)

type functionThatDoesNothing func (string)

// pretty normal
func div(num float32, denom float32) float32 {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// You can return multiple values
func divWithReminder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num/denom, num %denom, nil
}

// In Go we can do named return values. Basically, pre-declaring return values.
func divWithReminderNamed(num int, denom int) (result int, remainder int, e error) {
	if denom == 0 {
		e = errors.New("cannot divide by zero")
		return result, remainder, e
	}
	result = num/denom
	remainder = num % denom
	return result, remainder, e
}

func readOutLoud(name string) {
	fmt.Println("Ladies and Gentlemen, Introducing!!!!!", name)
}
	
func main() {
	fmt.Println("What's 10/42?", div(10, 42))
	calc, remainder, _ := divWithReminder(10, 42) // you cannot assign to a single variable.
	calcAlpha, remainderAlpha, _ := divWithReminderNamed(10, 42)
	fmt.Println("Let see the other one. Whats 10/42?", calc, remainder)
	fmt.Println("Like the previous one but with named return values!", calcAlpha, remainderAlpha)

	// Go is a functional programming language. Meaning you can do with functions about anything you can do in JS, Py

	// Assign a function to a value and call as a function
	funnyFunc := readOutLoud
	funnyFunc("Leonidas")

	var typedFunnyFunc functionThatDoesNothing = readOutLoud
	typedFunnyFunc("Leleonidas")

	// Here is an anonymous function
	// With this you can unlock everything else
	f := func (j int) {
		// count backwards
		for j > 0 {
			fmt.Println("Counting", j)
			j--
		}
	}

	f(10)
}
