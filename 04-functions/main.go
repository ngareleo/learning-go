package main

import "fmt"

func main() {
	fmt.Println("What's 10/42?", div(10, 42))
}

func div(num float32, denom float32) float32 {
	if denom == 0 {
		return 0
	}
	return float32(num) / float32(denom)
}