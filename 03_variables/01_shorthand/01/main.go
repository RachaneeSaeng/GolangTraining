package main

import "fmt"

func main() {
	//cc = iota --can use with const only
	a := 10
	b := "golang"
	c := 4.57
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	var h = a + int(c)

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
}
