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
		fmt.Println("You need at least an instagram account to enter")
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


	///////////////////////////////////////
	///////// For loop///// //////////////
	///////////////////////////////////////

	// For-loops are the only looping mechanisms in Go
	// It has several styles
	// 1. Traditional style
	for i := 0; i < 3; i++ {
		fmt.Printf("Leo you're amazing %d times\n", i)
	}
	//or
	i := 0
	for i < 3 {
		fmt.Printf("Leo you're very amazing x %d\n", i)
		i++
	}
	// 2. Infinite loop
	j := 0
	for {
		fmt.Println("Hello traveler")
		j++
		if j == 5 {
			break
		}
	}
	// 3. For range
	colors := [] string {"red", "blue", "green", "purple"}
	for i, v := range colors { // we get two variables i is the index and v is the value
		fmt.Println(i + 1, " Color ", v)
	}
	// incase you don't need the index
	for _, v := range colors {
		fmt.Println(v)
	}
	// it also works for maps, let say
	spells := map[string] int {
		"hate": 30,
		"ice": 23,
		"poison": 100,
	}
	for k, v := range spells {
		fmt.Printf("You can get %s for %d damage\n", k, v)
	}
	// but do note that the order of execution will always vary. This is a security feature in Go
	// You can use a for-range loop also to loop over strings
	for k := range spells {
		fmt.Println(k)
		for _, w := range k { // The for-range loop will convert the characters to runes not bytes
			fmt.Println(w, ", ", string(w))
		}
		fmt.Print("\n")
	}
	nums := [] int {1, 2, 3, 4, 5}
	// do note that each iteration is a `copy` and not reference, so
	for _, num := range nums {
		num += 10
		fmt.Println(num)
	}
	fmt.Println("But the nums will still be", nums)

	///////////////////////////////////////
	///////// Switch //////////////////////
	///////////////////////////////////////

	// Switch statements work very much like other languages 
	
	for _, color := range colors {
		switch size := len(color); size {
			case 0, 1:
				fmt.Println(color, "What?!?!")
			case 2, 3, 4, 5:
				fmt.Println(color ,"Kind of expected")
			default:
				fmt.Println(color, "What are these exaggerated complex colors?")

		}
	}
}