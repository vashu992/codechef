package main

import "fmt"

func canEnterToken(rollNumber int) string {
	if rollNumber == 6 {
		return "YES"
	}
	return "NO"
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var rollNumber int
		fmt.Scan(&rollNumber)
		result := canEnterToken(rollNumber)
		fmt.Println(result)
	}
}
