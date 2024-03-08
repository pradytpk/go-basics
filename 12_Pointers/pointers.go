package main

import (
	"fmt"
)

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}
func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

// --Summary:
//
//	Create a program that can activate and deactivate security tags
//	on products.
//
// --Requirements:
// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

// * Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}
func inactivate(tag *SecurityTag) {
	*tag = Inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("checking out")
	for i := 0; i < len(items); i++ {
		inactivate(&items[i].tag)
	}
}

// func main() {
// 	counter := Counter{}
// 	hello := "Hello"
// 	world := "World"
// 	fmt.Println(hello, world)
// 	replace(&hello, "hi", &counter)
// 	fmt.Println(hello, world)
// 	phrase := []string{hello, world}
// 	fmt.Println(phrase)
// 	replace(&phrase[1], "G0!", &counter)
// 	fmt.Println(phrase)
// 	//* Perform the following:
// 	//  - Create at least 4 items, all with active security tags
// 	shirt := Item{"Shirt", Active}
// 	pants := Item{"pants", Active}
// 	purse := Item{"purse", Active}
// 	watch := Item{"watch", Active}
// 	//  - Store them in a slice or array
// 	items := []Item{shirt, pants, purse, watch}
// 	fmt.Println("Intial", items)
// 	//  - Deactivate any one security tag in the array/slice
// 	inactivate(&items[0].tag)
// 	fmt.Println("Item 0 deactivated", items)
// 	//  - Call the checkout() function to deactivate all tags
// 	checkout(items)
// 	//  - Print out the array/slice after each change
// 	fmt.Println("Checkout", items)
// }
