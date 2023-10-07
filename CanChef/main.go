package main

import "fmt"

func main() {
	var T, X, Y int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X, &Y)
		totalDistance := 2 * Y

		if totalDistance <= X * 15 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
