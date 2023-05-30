package main

import "fmt"

func calculateMinimumInvestment(earnings, taxThreshold int) int {
	minInvestment := earnings - taxThreshold
	return minInvestment
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var earnings, taxThreshold int
		fmt.Scan(&earnings, &taxThreshold)
		minInvestment := calculateMinimumInvestment(earnings, taxThreshold)
		fmt.Println(minInvestment)
	}
}
