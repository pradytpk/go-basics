package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//--Summary:
//  Create an interactive command line application that supports arbitrary commands.
// When the user enters a command, the program will respond with a message.
// The program should keep track of how many commands have been ran, and how many lines of text was entered by the user.
//
const (
	Hello   = "hello"
	Goodbye = "goodbye"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error Reading stdin:", inputErr)
		}
	}
	fmt.Printf("sum:%v\n", sum)

	//--Requirements:
	//* When the user enters either "hello" or "bye", the program
	//  should respond with a message after pressing the enter key.

	//* Upon program exit, some usage statistics should be printed
	//  ('Q' and 'q' do not count towards these statistics):
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0
	for scanner.Scan() {
		//* Whenever the user types a "Q" or "q", the program should exit.
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())
			switch text {
			case Hello:
				numCommands += 1
				fmt.Println("commad response hi")
			case Goodbye:
				numCommands += 1
				fmt.Println("commad response bye")
			}
			if text != "" {
				numLines += 1
			}
		}
	}
	//  - The number of non-blank lines entered
	//  - The number of commands entered
	fmt.Printf("You entered %v lines\n", numLines)
	fmt.Printf("You entered %v commands", numCommands)
}

//--Notes
//* Use any Reader implementation from the stdlib to implement the program
