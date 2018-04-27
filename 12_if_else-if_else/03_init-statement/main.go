package main

import "fmt"

func main() {

	b := true

	if food, drink := "Chocolate", "Cola"; b {
		fmt.Println(food)
		fmt.Println(drink)
	}

}
