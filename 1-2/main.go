package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Program:\t%q\n", os.Args[0])
	for index, value := range os.Args[1:] {
		fmt.Printf("Index:\t%d\tValue: %v\n", index, value)
	}
}
