package main

import "fmt"

func main() {
	var x [256]int
	var addr = &x[0]
	x[0] = 10
	fmt.Println(addr)

	var y = x[0]
	var addr2 = &y
	y = 20
	fmt.Println(addr2)

	fmt.Println(x[0])
	fmt.Println(y)
}
