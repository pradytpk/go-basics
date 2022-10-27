package main

import (
	"fmt"
)

type Passenger struct {
	Name     string
	TicketNo int
	Boarded  bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	pradeep := Passenger{"Test", 1, true}
	fmt.Println(pradeep)
	var (
		bill = Passenger{Name: "bill", TicketNo: 2}
		ella = Passenger{Name: "ella", TicketNo: 3}
	)
	fmt.Println(bill, ella)

	var hei Passenger
	hei.Name = "test2"
	hei.TicketNo = 4
	fmt.Println(hei)
	bill.Boarded = true
	ella.Boarded = false
	if bill.Boarded {
		fmt.Println("has boarded the bus", bill.Name)
	} else if ella.Boarded {
		fmt.Println("has boarded the bus", ella.Name)
	}
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rect)
	//* After performing the above requirements, double the size of the existing rectangle and repeat the calculations
	rect.a.y *= 2
	rect.b.x *= 2
	//  Print the new results to the terminal
	printInfo(rect)
}

//--Summary:
//  Create a program to calculate the area and perimeter of a rectangle.

//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

//* Create a rectangle structure containing its coordinates
type Coordinate struct {
	x, y int
}
type Rectangle struct {
	a Coordinate
	b Coordinate
}

//* Using functions, calculate the area and perimeter of a rectangle,Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter

func width(rec Rectangle) int {
	return (rec.b.x - rec.a.x)
}
func length(rec Rectangle) int {
	return (rec.b.y - rec.a.y)
}
func area(rec Rectangle) int {
	return length(rec) * width(rec)
}
func perimeter(rec Rectangle) int {
	return (width(rec) * 2) + (length(rec) * 2)
}
func printInfo(rec Rectangle) {
	fmt.Println("Area is", area(rec))
	fmt.Println("Perimeter is", perimeter(rec))
}
