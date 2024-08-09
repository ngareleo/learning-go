package main

import (
	"fmt"
	"maps"
	"slices"
)

// Go runtime
// The runtime is responsible for memory allocation, garbage collection, network, IO etc
// Unlike most languages that depend execute on a virtual environment,
// When we compile go code, the compiler writes the go runtime into the binary
// So when you ship go binary, you don't need a go environment. Downside is that binaries become big at-least 2MB
func main() {

	/////////////////////////////////////////////////////////
	///////////// Arrays   ////////////////////////////////
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
	///////////// Handling arrays 
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
	///////////// Slices ///////////////////////////////////
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
	////////////// Capacity 
	// Length of a size provided by `len()` gives us the number of elements currently allocated in a slice
	// However, the go runtime assigns extra memory (to give room during cycles where a call to `append` 
	// doesn't need invocation of garbage collector and mem allocator to create more memory). The thread can just add the element
	// To get the capacity of a slice use `cap()`
	fmt.Printf("The length of h is %d but capacity is %d\n", len(h), cap(h))
	// But if you know the size of the slice ahead of time you can use `make()` to create a slice
	i := make([]int, 2) // creates a slice of capacity 10
	// However, we have a gotcha here. If we try to
	fmt.Println("Lets use append on I. We get ", append(i, 10)) // [0 0 10]
	// This is because append always increases the length of the slice
	// To achieve *creating a pre-allocated capacity slice then immediately populate it* you can do this
	p := make([]string, 0, 20) // 0 is the length and 20 is the capacity. 
	p = append(append(p, l...), "purple") // add elements of l then purple at the end
	fmt.Println("P is now ", p)
	////////////// Clear 
	// Calling `clear` sets all elements to their default value BUT keeps the length the same
	// It mutates the slice
	clear(p)
	fmt.Println("Now we don't have any colors 😞", p)
	fmt.Printf("But p has cap(%d) and len(%d)\n", cap(p), len(p))
	// Tip: Ideally, it better to wasted memory than wasting cycles recollecting and reallocating memory
	// So if you know ahead of time, that a slice will take up 1000 elements max, just allocate 1000 capacity
	// even if you add 40 elements
	// instead of slowly growing memory to 40, just take the cost but again, it depends on use-case

	// Slices from slices
	// They work pretty much like slices in python [start:end]
	q := l[:2] // First two elements of L
	fmt.Println("Slice of L ", q)
	// HOWEVER, q shares memory with the original L, so any change to q changes L. So be careful
	q[0] = "yellow"
	fmt.Println("You've changed L too. See! ", l) 
	// Rule of thumb to make things easier for self, do not use `append()` with sub-slices

	// To make copies that don't overwrite each other, you can use the `copy()` function

	r := make([]string, 5)
	copy(r, l) // Copy takes, (destination_slice, source_slice) and copies values from source to dest until one of them runs out of elements
	// It returns the number of elements copied. So you can ignore
	fmt.Println("R now looks like ", r)
	// and if I change r
	r[0] = "teal"
	fmt.Printf("R now is %d and L is still %d\n", r, l)

	// Take array d we created at the beginning, you can convert to slice by
	s := d[:] 
	// Now we can grow s
	s = append(s, 32)
	fmt.Println("We converted d to a slice and now we grew it to ", s) // not the array but the slice

	 
	//////////////////////////////////////////////////////////
	///////////// Strings ///////////////////////////////////
	//////////////////////////////////////////////////////////

	// You can fuvk around with strings the same way as slices in terms of indexing and slicing so:
	t := "don't stop"
	fmt.Println("Don't ..", t[6:])
	// conversions in Go play very much like in C like
	fmt.Println("D in 'don't' we get", int(t[0]))
	fmt.Println("We can get bytes in a string like ", []byte(t))


	//////////////////////////////////////////////////////////
	///////////// Maps /////////////////////////////////////
	/////////////////////////////////////////////////////////

	// In Go you declare maps like map[KeyType]ValueType
	var u map[string]int // U is a nill
	fmt.Println("U ", u)
	// You can also declare a map like:
	v := map[string]int{
		"red": 1,
		"blue": 2,
		"green": 3, // You must add a trailing comma
	}
	fmt.Println("V ", v)

	// If you know the size you expect ahead of time you could
	w := make(map[string][]int, 10)
	// Maps are like slices, in that 
	// 1. They grow in size
	// 2. `len` gives you the number of items (key-value) pairs
	// 3. Default value is `nill`
	w["purple"] = []int {255, 0, 255}
	w["red"] = []int {255, 0, 0}
	fmt.Println("w looks like ", w)
	// You can use this syntax called `Comma ok idiom` to tell the difference between assigned keys and nil values like so
	x, y := w["blue"] // x, is the value, y, is True is "purple exists" and vice-versa
	fmt.Printf("Value is %s, Was is assigned? %d\n", x, y)
	// You can delete a key by using `delete()` 
	// The first argument is the map, the second arg is the key
	delete(w, "purple")
	// You can also empty a map using `clear()`
	clear(w) // It deletes everything and sets the map length to 0

	// You can check equality using
	z := map[string]int{
		"red": 1,
		"green": 3,
		"blue": 2,
	}
	aa := maps.Equal(v, z) 
	fmt.Println("Are the two maps equal? ", aa)

	//////////////////////////////////////////////////////////
	///////////// Structs /////////////////////////////////////
	/////////////////////////////////////////////////////////

	type vehicle struct {
		year int
		wheels int
		model string
	}

	golf := vehicle {
		year: 2000,
		wheels: 4,
		model: "Volkswagen",
	}
	fmt.Println("Sample vehicle ", golf)
	// You can create anonymous structs on the fly
	user := struct {
		name string
		age int
	}{
		name: "Traveller Joe",
		age: 2000,
	}
	fmt.Println("Sample user ", user)
 }
