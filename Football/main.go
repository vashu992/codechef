package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)

		if x == y && (x > 0 || y > 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
