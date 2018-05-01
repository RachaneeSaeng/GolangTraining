package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Variables in Go are always mutable. When a string variable is changed,
	//the internal fields of the variable (pointer and length) are changed. The address of variable never changes.

	var x string = "abc"
	fmt.Println(x, &x, (*reflect.StringHeader)(unsafe.Pointer(&x)))
	x = "cde"
	fmt.Println(x, &x, (*reflect.StringHeader)(unsafe.Pointer(&x)))
	x = x[1:]
	fmt.Println(x, &x, (*reflect.StringHeader)(unsafe.Pointer(&x)))
}
