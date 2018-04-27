package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func main() {
	a := 5
	fmt.Printf("%p\n", &a) // address in main
	fmt.Println(&a)        // address in main
	zero(a)
	fmt.Println(a) // x is still 5
}
