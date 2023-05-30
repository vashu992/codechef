package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x, h int
		fmt.Scan(&x, &h)

		if x >= h {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
