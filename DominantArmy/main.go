package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var NA, NB, NC int
		fmt.Scan(&NA, &NB, &NC)

		if (NA > NB+NC) || (NB > NA+NC) || (NC > NA+NB) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
