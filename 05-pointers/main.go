package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	id int8
}

func sayMyNameLoudly (name *string) {
	if name != nil {
		upCaseName := strings.ToUpper(*name)
		fmt.Println("Ladies and gentlemen, introducing!!!!", upCaseName, "Clap everybody")
	}
}

// When we pass a pointer to a function, the pointer is copied
// Meaning, we get a copy of the memory address, not the pointer itself
// So when we pass a pointer, and you try to reassign the pointer to another value
// You just changed the copy, and the actual pointer outside the function still retains the same memory address
// (Personal take) It helps prevent allocating memory in a function and allocating it to a pointer outside that will out-live the function
// Plus its better to know that when you pass in a pointer after a function call, you'll still point to the same place
// the only side-effect is that the value of the pointer.
func convertNameToUppercase (person *Person) {
	if person != nil {
		person.name = strings.ToUpper(person.name)
	}
}

func main () {
	// Pointers in Go are pretty much like pointers everywhere else
	a := 10
	b := &a // B now holds memory address of a
	fmt.Println("What address is A in", b) 

	// we can declare b like this too
	var c *int // C now has a value of nil
	fmt.Println("What is the value of C?",c)

	// then we can reassign it later
	c = &a
	fmt.Println("C has same address as B now. Right? B and C", b, c)

	// Go however doesn't have pointer arithmetic like C
	// So we cannot use that magic to walk an array or sth

	// Constants can also not have a pointer because a pointer is a memory location
	// Constants are resolved during compile time and don't have memory address

	// We also have the indirection * to read the value of a pointer
	d := "Hello Traveler"
	e := &d
	fmt.Println("Now we can read the value of pointer E, ", *e, "from memory location", e)

	// If we try to dereference a nill pointer, Go runtime panics
	// So you need to check first. Like sayMyNameLoudly has done
	username := "mini-boss"
	sayMyNameLoudly(&username)

	miniboss := Person {
		name: "Mini boss",
		id: 2,
	}
	// This is now the power of pointers, we can mutate. 
	// Personally, no return type indicates better that we have mutated
	convertNameToUppercase(&miniboss)

	fmt.Println("What's your name boss? ", miniboss.name)


	// Here's a pointers caveat
	type Vehicle struct {
		make string
		owner *string
		year int
	}
	// golf := Vehicle {
	// 	make: "Golf",
	// 	owner: "Leo",
	// 	year: 2012,
	// }
	// Above assignment will panic, because we pass a value instead of pointer. To fix, we need a helper
	r := func (v string) *string {
		return &v
	}
	golf := Vehicle {
		make: "Golf",
		owner: r("Mini-boss"),
		year: 2012,
	}
	// To make it neater, you can use a factory method to do this if needed
	fmt.Println("Who owns this beautiful classic but modern car?", *golf.owner)


}