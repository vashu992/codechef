package main

import "fmt"

func main() {
	var T, X int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&X)

		if X < 3 {
			fmt.Println("LIGHT")
		} else if X >= 3 && X < 7 {
			fmt.Println("MODERATE")
		} else {
			fmt.Println("HEAVY")
		}
	}
}
