package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {

		var n int
		fmt.Scan(&n)

		totalWater := 2 * n
		fmt.Println(totalWater)
	}

}
