package main

import "fmt"

func auctionWinner(a, b, c int) string {
	if a > b && a > c {
		return "Alice"
	} else if b > a && b > c {
		return "Bob"
	} else {
		return "Charlie"
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		winner := auctionWinner(a, b, c)
		fmt.Println(winner)
	}
}
