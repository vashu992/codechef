package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		minCost := min(3*X, 2*Y)

		fmt.Println(minCost)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
