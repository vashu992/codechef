package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N, X int
		fmt.Scan(&N, &X)

		if X*2 >= N {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
