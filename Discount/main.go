package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x int
		fmt.Scan(&x)

		Discount := 100 - x
		fmt.Println(Discount)
	}
}
