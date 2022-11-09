package main

import (
	"fmt"
	"unicode"
)

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2+2=", compute(2, 2, add))
	fmt.Println("10-2=", compute(10, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))
	mul := func(lhs, rhs int) int {
		fmt.Printf("mulitiply %v *%v=", lhs, rhs)
		return lhs * rhs
	}
	fmt.Println(compute(3, 3, mul))
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	//* Using closures, determine the following information about the text and
	//  print a report to the terminal:
	//  - Number of letters
	//  - Number of digits
	//  - Number of spaces
	//  - Number of punctuation marks
	letters := 0
	numbers := 0
	spaces := 0
	punctuation := 0

	lineFunc := func(line string) {
		for _, v := range line {
			if unicode.IsLetter(v) {
				letters += 1
			}
			if unicode.IsNumber(v) {
				numbers += 1
			}
			if unicode.IsSpace(v) {
				spaces += 1
			}
			if unicode.IsPunct(v) {
				punctuation += 1
			}
		}
	}
	lineIterator(lines, lineFunc)
	fmt.Println(spaces, "spaces")
	fmt.Println(numbers, "numbers")
	fmt.Println(punctuation, "punctuation")
	fmt.Println(letters, "letters")
}

type LineCallBack func(line string)

// --Summary:
//
//	Create a program that can create a report of rune information from
//	lines of text.
//
// --Requirements:
//   - Create a single function to iterate over each line of text that is
//     provided in main().
func lineIterator(lines []string, callback LineCallBack) {
	for i := 0; i < len(lines); i++ {
		//  - The function must return nothing and must execute a closure
		callback(lines[i])
	}
}

//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification
