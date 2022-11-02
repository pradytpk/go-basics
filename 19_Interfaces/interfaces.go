package main

import "fmt"

type Preparer interface {
	PrepareDish()
}
type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--dish:%v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Chicken("wings"),
		Salad("peas"),
	}
	prepareDishes(dishes)

	//* Direct at least 1 of each vehicle type to the correct lift, and print out the vehicle information.
	car := Car("Sporty")
	truck := Truck("Monster")
	Motorcycle := Motorcycle("TVS")
	sendToLift(car)
	sendToLift(truck)
	sendToLift(Motorcycle)
}

//--Summary:
//  Create a program that directs vehicles at a mechanic shop to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

//* Write a single function to handle all of the vehicles that the shop works on.
type Motorcycle string
type Car string
type Truck string

//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle:%v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car:%v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck:%v", string(t))
}
func (m Motorcycle) PickLift() Lift {
	return SmallLift
}
func (c Car) PickLift() Lift {
	return StandardLift
}
func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to Small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to Standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to Large lift\n", p)
	}
}

//--Notes:
//* Use any names for vehicle models
