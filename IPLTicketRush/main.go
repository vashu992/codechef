package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		unableToBook := 0
		if n > m {
			unableToBook = n - m
		}

		fmt.Println(unableToBook)
	}
}
