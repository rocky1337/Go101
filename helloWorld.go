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

// TODO: Figure out how to run this in a running console app
// TODO: Handle multiple 2+2=4  4+5= 9
// TODO: Add a clear to restart to no numbers and what not.
func main() {
	// Inputs
	var total int
	wordPtr := flag.String("math", "foo", "a string")
	numPtr1 := flag.Int("num1", 0, "the First Int")
	numPtr2 := flag.Int("num2", 0, "the Second Int")
	// Must parse to get actual value of the pointers to the inputs
	flag.Parse()

	switch *wordPtr {
	case "add":
		total = add(*numPtr1, *numPtr2)
	case "subtract":
		total = subtract(*numPtr1, *numPtr2)
	case "multiply":
		total = multiply(*numPtr1, *numPtr2)
	case "divide":
		total = divide(*numPtr1, *numPtr2)
	default:
		fmt.Println("Incorrect input please reconsider your life choices")
	}

	fmt.Println(total)

}
