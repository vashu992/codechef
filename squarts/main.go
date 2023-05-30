package main

import "fmt"

func calculateSquats(sets int) int {
	squatsPerSet := 15
	totalSquats := sets * squatsPerSet
	return totalSquats
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var sets int
		fmt.Scan(&sets)
		totalSquats := calculateSquats(sets)
		fmt.Println(totalSquats)
	}
}
