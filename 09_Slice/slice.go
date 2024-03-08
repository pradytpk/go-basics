package main

import "fmt"

func main() {
	route := []string{"Pradeep", "test", "test2"}
	printSlice("Route 1", route)
	route = append(route, "Home")
	printSlice("Route 2", route)
	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")
	route = route[2:]
	printSlice("Remaining Locaitons", route)
}

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("_____", title, "_____")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}

	//* Perform the following:
	//  - Create an assembly line having any three part

	assemblyLine := make([]Part, 3)
	assemblyLine[0] = "pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "sheet"
	fmt.Println("\n3 Parts")
	printLine(assemblyLine)
	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "washer", "wheel")
	fmt.Println("\nadded 2 Parts")
	printLine(assemblyLine)
	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced")
	printLine(assemblyLine)
	//  - Print out the contents of the assembly line at each step

}

//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
type Part string

//* Create a function to print out the contents of the assembly line
func printLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts
