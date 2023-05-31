package main

import "fmt"

func main() {
	var r, c, e int
	fmt.Scan(&r, &c, &e)

	totalCells := (r + e) * c
	fmt.Println(totalCells)
}
