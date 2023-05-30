package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x, n int
		fmt.Scan(&x, &n)

		score := (x / 10) * n
		fmt.Println(score)
	}
}
