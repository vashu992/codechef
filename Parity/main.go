package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		if n%2 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
