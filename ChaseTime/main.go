package main

import "fmt"

func calculateMaxGames(hours int) int {
	return hours * 60 / 20
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var hours int
		fmt.Scan(&hours)

		maxGames := calculateMaxGames(hours)
		fmt.Println(maxGames)
	}
}
