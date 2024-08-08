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
	Literals()
	message := greetings.Hello("Leo")
	fmt.Println(message)

}


func Literals() {
	//////////////////////////////////////////////////
	//////////// Basics
	//////////////////////////////////////////////////

	var a bool = false // bool => boolean
	fmt.Println("A is %b", a)
	var b int = 32 // int is platform dependent. i.e on 32-bit machines
	fmt.Println("B is a platform dependent integer %d", b)
	// b is 32 bit
	// on 64-bit machines, b is 64 bit
	var c int32 = 43 // this is a 32 bit integer
	fmt.Println("C is a strictly 32-bit integer %d", c)
	// const sum = b + c;  This will panic why?
	// int being platform dependent, go compiler (can know ahead of time what's its compiling for) but choses
	// to refuse this operation to makes things easier
	var d byte = 32 // this is a unsigned 8-bit integer (uint8)
	fmt.Println("D is a unsigned 8-bit integer %d", d)


	//////////////////////////////////////////////////
	//////////// Type conversions
	//////////////////////////////////////////////////

	// Go's type conversion is simple
	// There's no implicit type conversion in Go
	// All type conversions are explicit
	var e int = 10
	f := float32(e) // F is now a floating point
	// The conversions functions are named the same as type 
	g := int32(f) // Back to a int32. You get the idea

	///////////////// NOTiCE ///////////////////////////

	// Go has no truthiness like Javascript where Boolean(0) is false or Boolean("") is also false
	// In fact you cannot explicitly convert a value to a bool unless using type comparison operators
	// So to convert 0 to a false value 
	h := 0
	y := h == 1 // Y is now false. You get the idea in JS i could have said y = Boolean(0) // true

	////////////////////////////////////////////////////////////
	///////////// So what's the difference between := and var?
	////////////////////////////////////////////////////////////

	var i int = 10
	j := 10
	// The only difference is that var (although highly discouraged) can be put in the global scope
	// While := is only restricted within functions
	// var is preferred in functions when you initialize but don't assign like
	var k bool // This is cleaner compared to k := bool()
	// Speaking of the above case, Go works with default values for types. Unassigned int is defaulted to 0
	// Unassigned string is defaulted to "". so 
	var l string // l := "" Same thing
}