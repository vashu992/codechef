package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ratedUsers := n - a
	highRatingUsers := ratedUsers - b

	fmt.Println(ratedUsers, highRatingUsers)
}
