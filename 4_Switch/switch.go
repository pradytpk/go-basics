package main

import "fmt"

//--Summary:
//  Create a program to display a classification based on age.
//--Requirements:

func price() int {
	return 10
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Print("cheap item")
	case p < 10:
		fmt.Print("Moderately priced item")
	default:
		fmt.Println("Expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy Seating")
	case Business:
		fmt.Println("Business Seating")
	case FirstClass:
		fmt.Println("FirstClass Seating")
	default:
		fmt.Println("Other Seating")
	}
	//* Use a `switch` statement to print the following:

	switch age := 0; {
	//  - "newborn" when age is 0
	case age == 0:
		fmt.Println("newborn")
	//  - "toddler" when age is 1, 2, or 3
	case age >= 1 && age <= 3:
		fmt.Println("toddler")
		//  - "child" when age is 4 through 12
	case age < 13:
		fmt.Println("child")
		//  - "teenager" when age is 13 through 17
	case age < 18:
		fmt.Println("teenager")
		//  - "adult" when age is 18+
	default:
		fmt.Println("teenager")
	}
}
