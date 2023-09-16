package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T) 

	for t := 0; t < T; t++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		
		if Y >= X*3 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
