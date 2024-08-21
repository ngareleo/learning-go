package main

import (
	"fmt"
	"learninggo/modules/math"
)

func main() {
	p1 := math.NewPoint(2, 3)
	p2 := math.NewPoint(4, 5)
	p1.Attach(&p2)
	fmt.Println("What's the difference between p1 and p2?", p1.Distance())
}
