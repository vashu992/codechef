package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var W, X, Y, Z int
		fmt.Scan(&W, &X, &Y, &Z)

		totalInflow := W + (Y * Z)

		if totalInflow > X {
			fmt.Println("overflow")
		} else if totalInflow == X {
			fmt.Println("filled")

		} else {
			fmt.Println("unfilled")
		}
	}
}
