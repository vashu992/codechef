package main

import "fmt"

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m)
	fmt.Scan(&x, &y)

	bones := n * x
	blood := m * y

	total := bones + blood

	fmt.Println(total)

}
