package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		limit := 1.07 * float64(X)

		if float64(Y) <= limit {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
