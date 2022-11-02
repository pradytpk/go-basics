package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index:%v", index))
	} else {
		return s.values[index], nil
	}
}
func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}

//--Summary:
//  Create a function that can parse time strings into component values.
//

type Time struct {
	hour, minute, second int
}
type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v:%v", t.msg, t.input)
}

//* Example time string: 14:07:33
func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error Prasing hour:%v", err), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error Prasing minute:%v", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error Prasing second:%v", err), input}
		}
		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hour out of range:0<=hour <=23", fmt.Sprintf("%v", hour)}
		}
		if minute > 59 || minute < 0 {
			return Time{}, &TimeParseError{"minute out of range:0<=minute <=23", fmt.Sprintf("%v", minute)}
		}
		if second > 59 || second < 0 {
			return Time{}, &TimeParseError{"minute out of range:0<=second <=23", fmt.Sprintf("%v", second)}
		}
		return Time{minute, hour, second}, nil
	}
}

//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:

//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors
