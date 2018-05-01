package main

import "fmt"

// set of const increment by 1
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
	x  = 1 >> 3
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t%d\n", KB, KB) // 1024 : 1 (0000000000000001) shift left 10 bit (0000010000000000 = 2^10 = 1024)
	//fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
	fmt.Printf("%b\t%d\n", x, x) // 0 : cannot rotate

	fmt.Printf("%T\t%b\n", 5, 5)
	fmt.Println(^5) // -6 : two's complement
	fmt.Printf("%T\t%b\n", -6, -6)
	fmt.Println(3 &^ 6)
}
