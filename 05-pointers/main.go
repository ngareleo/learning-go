package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	name string
	id   int8
}

func sayMyNameLoudly(name *string) {
	if name != nil {
		upCaseName := strings.ToUpper(*name)
		fmt.Println("Ladies and gentlemen, introducing!!!", upCaseName, "Clap everybody!!!")
	}
}

func multiplySecondValue(arr []int) {
	if len(arr) >= 2 {
		arr[1] *= 2
	}
}

// This is pretty bad, haven't yet dealt with IO but excuse me, trying to put a point across
func capitalizeNamesFromFile(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Something went wrong ", err)
		return
	}
	defer file.Close()
	wordCount := 0
	// The point is, don't litter please!
	// Whenever you are dealing with pointers, don't create them haphazardly, and leave them
	// Instead, put them together because it makes it easier for the garbage collector to collect them
	// It's called mechanical sympathy. We know that reading a continuous location in memory is faster than reading as many
	// scattered memory locations

	// The Go Garbage collector (GC) runs whenever the heap size reaches a set-maximum size
	// So whenever it runs, the quicker it can read from memory the faster the GC will finish and the faster
	// the thread will go back to execution
	data := make([]byte, 100)
	names := make([]string, 100)
	for {
		_, err = file.Read(data)
		for _, v := range data {
			if rune(v) == '\n' {
				names = append(names, "")
				wordCount++
			} else {
				names[wordCount] += string(v)
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("reached EOF")
			}
			fmt.Println("Names found ", names)
			return
		}
	}

}

// When we pass a pointer to a function, the pointer is copied
// Meaning, we get a copy of the memory address, not the pointer itself
// So when we pass a pointer, and you try to reassign the pointer to another value
// You just changed the copy, and the actual pointer outside the function still retains the same memory address
// (Personal take) It helps prevent allocating memory in a function and allocating it to a pointer outside that will out-live the function
// Plus its better to know that when you pass in a pointer after a function call, you'll still point to the same place
// the only side-effect is that the value of the pointer.
func convertNameToUppercase(person *Person) {
	if person != nil {
		person.name = strings.ToUpper(person.name)
	}
}

func main() {
	// Pointers in Go are pretty much like pointers everywhere else
	a := 10
	b := &a // B now holds memory address of a
	fmt.Println("What address is A in", b)

	// we can declare b like this too
	var c *int // C now has a value of nil
	fmt.Println("What is the value of C?", c)

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

	miniboss := Person{
		name: "Mini boss",
		id:   2,
	}
	// This is now the power of pointers, we can mutate.
	// Personally, no return type indicates better that we have mutated
	convertNameToUppercase(&miniboss)

	fmt.Println("What's your name boss? ", miniboss.name)

	// Here's a pointers caveat
	type Vehicle struct {
		make  string
		owner *string
		year  int
	}
	// golf := Vehicle {
	// 	make: "Golf",
	// 	owner: "Leo",
	// 	year: 2012,
	// }
	// Above assignment will panic, because we pass a value instead of pointer. To fix, we need a helper
	r := func(v string) *string {
		return &v
	}
	golf := Vehicle{
		make:  "Golf",
		owner: r("Mini-boss"),
		year:  2012,
	}
	// To make it neater, you can use a factory method to do this if needed
	fmt.Println("Who owns this beautiful classic but modern car?", *golf.owner)

	// As for maps and slices, these are implemented using pointers
	// So if you pass a map as a function argument, you pass the pointer not a copy
	// That's why they can be mutated
	// Generally, the advice is so stay away from maps and use structs as there's nothing you gain from using
	// maps that you can't gain from using structs
	// The only exception is that maps are data-structures that allow you to store key-values but you don't
	// know the keys during *compile time*

	// Slices are a little bit different. They are structs with three fields, length, capacity and pointer to memory
	// When you pass it as an argument to a function, copies of all three are made
	// Making a change to the contents reflects back to the original because you're directly addressing the memory
	// But when you append, you're changing the length of the copy not the original and the changes do not reflect
	second := 3
	var g = []int{2, second, 4, 5}
	multiplySecondValue(g)
	fmt.Println("Did we mutate G?", g[1] == second*2, "G => ", g)

	capitalizeNamesFromFile("./sample_data.txt")
}
