package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var C, X, Y int
		fmt.Scan(&C, &X, &Y)

		requiredChocolate := C - X

		totalCost := requiredChocolate * Y

		fmt.Println(totalCost)
	}
}
