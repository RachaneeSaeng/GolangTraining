package main

import (
	"fmt"
)

// Weekday declare a new type named Weekday which will unify our enum values
// It has an underlying type of unsigned integer (uint).
type Weekday int

// Declare typed constants each with type of Weekday
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String returns the name of the day
func (day Weekday) String() string {

	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	// create arrya of string
	names := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return names[day]
}

// Weekend return true for a weekend day
func (day Weekday) Weekend() bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}

func main() {
	// Which day it is? Sunday
	fmt.Printf("Which day it is? %s\n", Sunday)

	// Is Saturday a weekend day? true
	fmt.Printf("Is Saturday a weekend day? %t\n", Saturday.Weekend())

	// Is Monday a weekend day? false
	fmt.Printf("Is Monday a weekend day? %t\n", Monday.Weekend())
}
