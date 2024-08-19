package main

import "fmt"

type Vehicle struct {
	model string
	year int
	topSpeed int
	speed int
	speedHistory []int
}

// Below is a method (not function because we've bound it to a type)
// Go allows you to declare methods only at the package level
// It also restricts method declaration within the same file as the type declaration 
// Meaning you can only bind types that you own
// (v Vehicle) is a receiver specification. This is the type that will be bound to a function
func (v Vehicle) String() string {
	return fmt.Sprintf("Model %s of year %d with a top speed of %d", v.model, v.year, v.topSpeed)
}

// You can also define using pointer type receivers
// In this case though, it is recommended to have consistency to avoid confusion around your types
// If you pass a value specification, use value specification all through and the vice-versa
func (v *Vehicle) Accelerate(by int) {
	v.speed += by
	v.speedHistory = append(v.speedHistory, v.speed)
}

func (v *Vehicle) fakeHistory() {
	var fake = []int { 50, 60, 73, 119, 120}

	// Best pattern is to code for nil values
	if v == nil {
		v = &Vehicle { speedHistory: []int{} }
	}

	v.speedHistory = append(v.speedHistory, fake...)
}

func readFakeSpeedHistory(v Vehicle) {
	// we call a pointer receiver method on a value function argument
	v.fakeHistory()
	for i, v := range(v.speedHistory) {
		fmt.Printf("[record %d] %dkm/h\n", i + 1, v)
	}
}


func main () {
	v := Vehicle {
		model: "Golf",
		year: 2017,
		topSpeed: 220,
		speed: 0,
	}
	fmt.Println(v) 
	// or
	fmt.Println(v.String()) // less idiomatic

	// Even if v is a value type, we were able to call a pointer receiver method
	// Go automatically takes the mem address of v so below we get (&v).Accelerate()
	v.Accelerate(10)
	fmt.Println("Moving at", v.speed)

	////////////////// WARNING ///////////////////////////////////////////
	// If you call a value receiver method on a nil pointer instance,   //
	// the code will compile but will panic during runtime              //
	////////////////// WARNING ///////////////////////////////////////////

	// Also function calls apply as usual
	// If you call a pointer receiver method for a value parameter in a function, you mutate the copy
	// not the original. go just copies the mem location of the copy
	// A good way of inserting temporary behavior

	// this function calls a pointer receiver method on v that edits speed history
	readFakeSpeedHistory(v)
	fmt.Println("Has the history changed?", v.speedHistory) // nope

	var polo Vehicle;
	// Go allow you to call methods on nil instances of a certain type
	polo.fakeHistory()
	fmt.Println("Polo history", polo.speedHistory)
}	