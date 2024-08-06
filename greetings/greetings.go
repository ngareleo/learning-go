package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hello captain %v", name)
	return message
}