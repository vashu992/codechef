package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)

		cost1 := a - c
		cost2 := b - d

		if cost1 < cost2 {
			fmt.Println("FIRST")
		} else if cost2 < cost1 {
			fmt.Println("SECOND")
		} else {
			fmt.Println("ANY")
		}
	}
}
