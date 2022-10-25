package main

import "fmt"

func average(a, b, c int) float32 {
	return float32(a+b+c) / 3
}

//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func weekday(day int) bool {
	return day <= 4
}

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8
	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 scored higher than quiz1")
	} else {
		fmt.Println("quiz1 & quiz2 have the same score1")
	}
	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Good")
	} else {
		fmt.Println("bad")
	}
	// The day and role. Change these to check your work.
	today, role := Monday, Guest
	//* Use the accessGranted() and accessDenied() functions to display informational messages

	if role == Admin || role == Manager {
		//* Access at any time: Admin, Manager
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		//* Access weekends: Contractor
		accessGranted()
	} else if role == Member && weekday(today) {
		//* Access weekdays: Member
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		//* Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
	} else {
		accessDenied()
	}

}
