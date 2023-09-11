package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, X, Y int
		fmt.Scan(&N, &X, &Y)

		if (X * Y) >= N {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
