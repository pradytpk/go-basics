package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}
func main() {
	rooms := [...]Room{
		{name: "office"},
		{name: "Reception"},
		{name: "ops"},
		{name: "Warehouse"},
	}
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness((rooms))
	//* Insert 3 products into the array
	shoppingList := [4]Product{
		{1, "Banana"},
		{3, "Apple"},
		{6, "Coconut"},
	}
	printStats(shoppingList)
	//* Add a fourth product to the list and print out the information again array, create a shopping list with enough room for 4 products
	shoppingList[3] = Product{4, "cherry"}
	printStats(shoppingList)
}

//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.

//Products must include the price and the name
type Product struct {
	price int
	name  string
}

// Using an number of items The total cost of the items
func printStats(list [4]Product) {
	var cost, totalItems int
	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price
		if item.name != "" {
			totalItems += 1
		}
	}
	//* Print to the terminal:
	//  - The last item on the list
	//  - The total
	fmt.Println("last item on the list:", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost", cost)
}
