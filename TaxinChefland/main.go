package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var income int
		fmt.Scan(&income)

		if income > 100 {
			income -= 10
		}

		fmt.Println(income)
	}
}
