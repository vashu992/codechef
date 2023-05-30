package main

import "fmt"

func checkTopTen(rank int) string {
	if rank <= 10 {
		return "YES"
	}
	return "NO"
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var rank int
		fmt.Scan(&rank)
		result := checkTopTen(rank)
		fmt.Println(result)
	}
}
