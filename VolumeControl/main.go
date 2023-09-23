package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		diff := abs(Y - X)

		fmt.Println(diff)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
