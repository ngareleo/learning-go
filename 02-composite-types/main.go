package main

import (
	"fmt"
	"slices"
)

func main() {
	/////////////////////////////////////////////////////////
	///////////// Arrays
	/////////////////////////////////////////////////////////
	var a [3] int // We have created a list of integers called x
	// By default all three elements are {0, 0, 0} due to Go's default assignment
	var b  = [3] int {0, 0, 0} // b and a have the same value
	// actually if you were to do 
	c := a == b // C will be True. Go compares arrays by length and by value of elements
	fmt.Printf("So are a and b equal? %f\n", c) // told you
	// To assign values, we do the old
	a[0] = 20
	a[1] = 21
	a[2] = 22
	fmt.Println("The array A is now ", a)
	// If you don't want to specify the size
	var d = [...] int {1, 2, 3, 4, 5} // a little bit too verbose
	var e [2][3] int // We have ourselves a multi-dim array
	fmt.Println("E ", e)
	///////////// Handling arrays ///////////////////////////////////
	fmt.Println("Can't count, here is length of d ", len(d))
	// Here's a kicker though
	// Arrays are rarely used in Go because of their limitations
	// The size of the array is included as the array's type. So a has type [3] int and d has [5] int
	// So why is that bad?
	// 1. You cannot use variables to declare arrays sizes, because types are resolved during compile time
	// 2. You cannot convert between arrays of different sizes because they are of different type
	// 3. In function declarations, you cannot just declare an array of any size as a parameter
	// So only use arrays when you know the size ahead of time 
	

	/////////////////////////////////////////////////////////
	///////////// Slices
	/////////////////////////////////////////////////////////

	// Slices are like arrays in that they can hold multiple values but can be of variable sizes
	// So why arrays and slices? Arrays are stores and slices are containers
	// Think of python slices which are segments of a larger array

	var f = [] int {1, 2, 3} // You don't need to declare the size of the array because it's not part of their type
	fmt.Println("Our first slice f-alpha-001 ", f)
	// To make it flexible Go compiler during compile time, just needs the necessary information to resolve types
	// You can create a staggered slice like this
	var g = [] int {1, 5: 4, 7, 8: 15, 2} // 5: 4 means, fill from index 1 to index 4 with 0's then at index [5] insert a 4, same for 8:15
	fmt.Println("I bet G looks weird ", g) // [1 0 0 0 0 4 7 0 15 2]
	var h [] int // Here we see the real difference between arrays and slices
	fmt.Println("I know you expect some default value like [0] but h is actually ", h) // nill. It's like null but they claim is "different"
	// You cannot compare slices like arrays using == operator
	// i := h == g This will panic because you can only compare slices to nill using == (Empty slices [or something])
	j := h == nil
	fmt.Println("So is it nill?", j)
	///////////// Handling slices ///////////////////////////////////	
	/////////////////// Compare a slice as long as the elements are comparable using `Equal`
	k := slices.Equal(f, g)
	fmt.Println("Are {f & g} they the same?", k)
	var l = [] string {"red", "green", "blue"}
	// fmt.Println(slices.Equal(g, l)) We cannot compare the two (Strings and Ints)
	////////////////// Get sizes 
	m := len(l)
	n := len(h)
	fmt.Printf("Sizes of l {%d} and h {%d}\n", m, n) // H is nill but we get sizes 0
	///////////////// Grow sizes
	// o := append(l, "orange") // It a pure function. 
	// To avoid growing mem
	l = append(l, "orange")
	fmt.Println("L now is ", l)
	// I can fill the h (its nil) properly like this
	h = append(h, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Voila H ", h)
	// I can merge h with f like so
	h = append(h, f...)
	fmt.Println("Now h has grown to ", h)
 }
