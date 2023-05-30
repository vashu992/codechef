package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		maxPeople := (n * 5) + (m * 7)
		fmt.Println(maxPeople)
	}
}
