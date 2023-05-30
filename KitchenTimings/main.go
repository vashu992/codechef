package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		hours := (y - x + 12) % 12
		fmt.Println(hours)
	}
}
