package main

import "fmt"

func main() {
	var T, X, Y int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X, &Y)

		chocolateTastiness := 2 * X
		candyTastiness := 5 * Y

		if chocolateTastiness > candyTastiness {
			fmt.Println("Chocolate")
		} else if chocolateTastiness < candyTastiness {
			fmt.Println("Candy")
		} else {
			fmt.Println("Either")
		}
	}
}
