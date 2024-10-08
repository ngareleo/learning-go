package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
)

func calcDivision(a, b int) (exp int, rem int, err error) {
	if b == 0 {
		// Ideally, you should return zero values when we return an error
		return exp, rem, errors.New("cannot divide by zero")
	}
	return a / b, a % b, nil
}

func madeToPanic() {
	// A panic is like throwing errors in many languages
	// A panic can be raised by the Go runtime, but you can also call panic() to exit immediately
	// The moment that a program panics the executing function will stop further execution, and will execute all the
	// defer functions is order
	// Then it will go up the call stack and execute defer too, until it reaches main, where it will call the defer
	// function there too and quit
	defer func() {
		// You can listen in on a panic by using
		if v := recover(); v != nil {
			fmt.Println("You stopped this func")
		}
	}()

	// The general advice is not to use this pattern for exception handling. Instead, use method below
	// Panic is reserved for cases where the program cannot recover from
	// recover() then serves as a means to gracefully shutdown
	panic("I'm made to panic")
}

// defining error states

type Status int

const (
	InvalidKey Status = iota
	UserNotFound
)

type StatusErr struct {
	message string
	status  Status
}

// To define a struct as an error.
// It needs to implement the Error interface by implementing an Error() function.
func (s StatusErr) Error() string {
	return s.message
}

// We need to implement this method to be able to compare errors using errors.Is(a, b)
func (s StatusErr) Is(err error) bool {
	statusErr := err.(StatusErr)
	return statusErr.status == s.status
}

func findUserByUsername(user string) (string, error) {
	return "", errors.New("cannot find user")
}

func checkPassword(pass string) (string, error) {
	return "", errors.New("incorrect password")
}

func grantToken(user string) string {
	return "dasdsvsdvsewefewf"
}

func LoginUser(user, pass string) (string, error) {

	_, err := findUserByUsername(user)
	if err != nil {
		return "", StatusErr{
			message: "cannot find user",
			status:  UserNotFound,
		}
	}
	_, passErr := checkPassword(pass)
	if passErr != nil {
		return "", StatusErr{
			message: "password incorrect",
			status:  InvalidKey,
		}
	}
	return grantToken(user), nil
}

func main() {

	_, _, err := calcDivision(3, 0)
	if err != nil {
		fmt.Println("This is how error handling is done", err)
	}

	data := []byte("Some random statement")
	notZipFl := bytes.NewReader(data)

	_, err2 := zip.NewReader(notZipFl, int64(len(data)))
	if err2 == zip.ErrFormat {
		fmt.Println("Should panic")
	}

	_, err3 := LoginUser("Leo", "12345")
	if errors.Is(err3, StatusErr{status: UserNotFound}) {
		fmt.Println("Couldn't find user")
	}

	madeToPanic()
}
