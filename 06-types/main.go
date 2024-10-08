package main

import "fmt"

// types are executable documentation. This is a common pattern in Go
// Use types to describe data and what we expect
type Year int
type Speed int

// This is a bad example but imagine this was something else with more complex behavior
// and edge cases, or could be computed using a complex formula
func (s Speed) compare(s2 Speed) bool {
	return s2 > s
}

func (s Speed) Float32() float32 {
	return float32(3)
}

// By declaring speed, we tie together the concept of speed for (topSpeed, speed, speedHistory) Vehicle
type Vehicle struct {
	model        string
	year         Year
	topSpeed     Speed
	speed        Speed
	speedHistory []Speed
}

// Below is a method (not function because we've bound it to a type)
// Go allows you to declare methods only at the package level
// It also restricts method declaration within the same file as the type declaration
// Meaning you can only bind types that you own
// (v Vehicle) is a receiver specification. This is the type that will be bound to a function
func (v Vehicle) String() string {
	return fmt.Sprintf("Model %s of year %d with a top speed of %d", v.model, v.year, v.topSpeed)
}

// You can also define using pointer type receivers
// In this case though, it is recommended to have consistency to avoid confusion around your types
// If you pass a value specification, use value specification all through and the vice-versa
func (v *Vehicle) Accelerate(by Speed) {
	if !(v.speed + by).compare(v.topSpeed) {
		v.speed += by
		v.speedHistory = append(v.speedHistory, v.speed)
	}
}

func (v *Vehicle) fakeHistory() {
	var fake = []Speed{50, 60, 73, 119, 120}

	// Best pattern is to code for nil values
	if v == nil {
		v = &Vehicle{speedHistory: []Speed{}}
	}

	v.speedHistory = append(v.speedHistory, fake...)
}

func readFakeSpeedHistory(v Vehicle) {
	// we call a pointer receiver method on a value function argument
	v.fakeHistory()
	for i, v := range v.speedHistory {
		fmt.Printf("[record %d] %d km/h\n", i+1, v)
	}
}

//////////////////////////////////////////////////////////////////////
//                     IOTA                                         //
//////////////////////////////////////////////////////////////////////

type DriveTrain int

const (
	aWD DriveTrain = iota
	fourWD
	rearWD
	frontWD
)

func (d DriveTrain) String() string {
	var repr string
	switch d {
	case aWD:
		repr = "All"
	case fourWD:
		repr = "Four wheel drive"
	case rearWD:
		repr = "Rear wheel drive"
	case frontWD:
		repr = "Front wheel drive"
	}
	return repr
}

//////////////////////////////////////////////////////////////////////
//                     Embeddings                                   //
//////////////////////////////////////////////////////////////////////

// Go prefers composition to inheritance because there's none
// So how do you share code?
// [[[[EMBEDDINGS]]]]]
// Truck now has properties of Vehicle and you can access them directly. ie. truck.model
// However, the two types are not interchangeable. So if a function declares function of type Vehicle, you cannot pass
// in Truck; but you can pass in truck.Vehicle
type Truck struct {
	Vehicle
	driveTrain      DriveTrain
	maxAcceleration Speed
}

func (t Truck) Describe() {
	fmt.Println("A truck with", t.driveTrain)
}

// If we create a method in the outer type that shadows one in the embedding, no voodoo is done
// To call this on an instance truck, call as truck.Accelerate(Speed(30)). To call the Accelerate method on the
func (t Truck) Accelerate(by Speed) {
	acc := 0
	if by > t.maxAcceleration {
		acc = int(t.maxAcceleration)
	} else {
		acc = int(by)
	}
	t.speed = Speed(int(t.speed) + acc)
}

func (t Truck) Decelerate(by Speed) {
	minSpeed := 0
	t.speed -= by

	if t.speed < Speed(minSpeed) {
		t.speed = Speed(minSpeed)
	}
}

func (t Truck) Speed () Speed {
	return t.speed
}
//////////////////////////////////////////////////////////////////////
//                     Interfaces                                   //
//////////////////////////////////////////////////////////////////////

type Locomotor interface {
	Speed() Speed
	Accelerate(Speed)
	Decelerate(Speed)
}


const pitLaneLimit = 30
func pitLaneOverride(l Locomotor) {
	l.Decelerate(l.Speed() - pitLaneLimit)
}

func main() {
	v := Vehicle{
		model:    "Golf",
		year:     2017,
		topSpeed: 220,
		speed:    0,
	}
	fmt.Println(v)
	// or
	fmt.Println(v.String()) // less idiomatic

	// Even if v is a value type, we were able to call a pointer receiver method
	// Go automatically takes the mem address of v so below we get (&v).Accelerate()
	v.Accelerate(10)
	fmt.Println("Moving at", v.speed)

	//////////////////////// [WARNING] ///////////////////////////////////
	//                                                                  //
	// If you call a value receiver method on a nil pointer instance,   //
	// the code will compile but will panic during runtime              //
	//////////////////////////////////////////////////////////////////////

	// Also function calls apply as usual
	// If you call a pointer receiver method for a value parameter in a function, you mutate the copy
	// not the original. go just copies the mem location of the copy
	// A good way of inserting temporary behavior

	// this function calls a pointer receiver method on v that edits speed history
	readFakeSpeedHistory(v)
	fmt.Println("Has the history changed?", v.speedHistory) // nope

	var polo Vehicle
	// Go allow you to call methods on nil instances of a certain type
	polo.fakeHistory()
	fmt.Println("Polo history", polo.speedHistory)

	truck := Truck{
		Vehicle: Vehicle{
			model: "Range Rover",
			year:  2010,
		},
		driveTrain:      aWD,
		maxAcceleration: 20,
	}

	// because of embeddings we can say
	truck.Accelerate(40)
	fmt.Printf("The %s is moving at %d km/h\n", truck.model, truck.speed)
	fmt.Println("Checking vehicle ", truck.Vehicle.model)

	pitLaneOverride(truck)
}
