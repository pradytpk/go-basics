package main

import (
	"fmt"
)

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1
	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)
	delete(shoppingList, "milk")
	fmt.Println("milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("need cereal?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yep", cereal, "boxes")
	}
	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("there are", totalItems, "on the shopping list")

	//* Create a map using the server names as the key and the server status as the value
	//* Set all of the server statuses to `Online` when creating the map
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}
	//  - call display server info function
	printServerStatus(serverStatus)
	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline
	//  - call display server info function
	printServerStatus(serverStatus)
	//  - change server status of all servers to `Maintenance`
	for server := range serverStatus {
		serverStatus[server] = Maintenance
	}
	//  - call display server info function
}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
func printServerStatus(servers map[string]int) {
	fmt.Println("\nthere are", len(servers), "Servers")
	//  - number of servers
	stas := make(map[int]int)
	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
	for _, status := range servers {
		switch status {
		case Online:
			stas[Online] += 1
		case Offline:
			stas[Offline] += 1
		case Maintenance:
			stas[Maintenance] += 1
		case Retired:
			stas[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}
	fmt.Println(stas[Online], "servers are online")
	fmt.Println(stas[Offline], "servers are Offline")
	fmt.Println(stas[Maintenance], "servers are Maintenance")
	fmt.Println(stas[Retired], "servers are Retired")
}

//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:

//* After creating the map, perform the following actions:
