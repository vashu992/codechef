package main

import "fmt"

func calculateFine(speed int) int {
	if speed <= 70 {
		return 0
	} else if speed > 70 && speed <= 100 {
		return 500
	} else {
		return 2000
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var speed int
		fmt.Scan(&speed)

		fine := calculateFine(speed)
		fmt.Println(fine)
	}
}
