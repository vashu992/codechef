package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	unAttempted := x - y
	fmt.Println(unAttempted)
}
