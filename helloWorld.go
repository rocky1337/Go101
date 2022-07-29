package main

import (
	"flag"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func main() {
	wordPtr := flag.String("math", "foo", "a string")

	fmt.Println(*wordPtr)

}
