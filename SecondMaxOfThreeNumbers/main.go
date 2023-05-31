package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		// Create a slice containing the three numbers
		numbers := []int{a, b, c}

		// Sort the numbers in ascending order
		sort.Ints(numbers)

		// Output the second-maximum number
		fmt.Println(numbers[1])
	}
}
