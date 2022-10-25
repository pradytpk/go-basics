package main

import (
	"fmt"
)

func main() {
	var myName = "pradeep"
	fmt.Println("My name is", myName)
	fmt.Println("My name is", myName, myName)

	var name string = "kumar"
	fmt.Println("name=", name)

	userName := "admin"
	fmt.Println("username=", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 4, 5
	fmt.Println("Part 1 is", part1, "other is", other)

	part2, other := 1, 0
	fmt.Println("Part 1 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

	//* Store your favorite color in a variable using the `var` keyword
	var favColor = "green"
	fmt.Println("My favorite color is", favColor)
	//* Store your birth year and age (in years) in two variables using compound assignment
	birthYears, ageInYears := 1990, 30
	fmt.Println("birth year is", birthYears, "age is", ageInYears)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstIntial = "P"
		lastIntial  = "K"
	)
	fmt.Println("First Intial is", firstIntial, "Last Intial is", lastIntial)
	//* Declare (but don't assign!) a variable for your age (in days), then assign it on the next line by multiplying 365 with the age variable created earlier
	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("I am", ageInDays, "age old")
}
