package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var X, Y, Z int
		fmt.Scan(&X, &Y, &Z)

		credits := (4 * X) + (2 * Y) + (0 * Z)
		fmt.Println(credits)
	}
}
