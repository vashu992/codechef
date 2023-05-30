package main

import "fmt"

func checkWaterIntake(water int) string {
	if water >= 2000 {
		return "YES"
	}
	return "NO"
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var water int
		fmt.Scan(&water)
		result := checkWaterIntake(water)
		fmt.Println(result)
	}
}
