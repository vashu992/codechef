package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N, X int
		fmt.Scan(&N, &X)

	
		songDuration := 3 * X

		timesListened := N / songDuration

		fmt.Println(timesListened)
	}
}
