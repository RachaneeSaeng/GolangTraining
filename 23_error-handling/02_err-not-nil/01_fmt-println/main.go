package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		log.Fatalln(err)
		panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
