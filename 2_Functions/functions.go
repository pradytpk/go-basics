package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

//--Summary:
//  Use functions to perform basic operations and print some information to the terminal.

//--Requirements:
//* Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello", name)
}

//* Write a function that returns any two numbers
func twoNo() (int, int) {
	return 3, 3
}

//* Write a function that returns any message, and call it from within fmt.Println()
func welcome() string {
	return "Hi pradeep"
}

//* Write a function to add 3 numbers together, supplied as arguments, and return the answer

func addThreeNo(a, b, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func anyNo() int {
	return 6
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's Dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 2)
	fmt.Println("A another baker's Dozen is", anotherBakersDozen)

	//* Call every function at least once
	greetPerson("Test1")
	fmt.Println(welcome())
	//* Add three numbers together using any combination of the existing functions.Print the result
	a, b := twoNo()
	answer := addThreeNo(anyNo(), a, b)
	fmt.Println(answer)
}
