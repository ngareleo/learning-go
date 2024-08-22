package main

import (
	"fmt"
	"learninggo/stdlib/io"
	"strings"
	"time"
)


func main() {
	s := "I love that Mary Jane!"
	vals, err := io.CountLetters(strings.NewReader(s))
	if err != nil {
		fmt.Println("Got error", err)
	}
	fmt.Println("What's the count?", vals)
)
}