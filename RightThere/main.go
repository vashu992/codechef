package main

import "fmt"

func canHostParty(N, X int) string {
	if N <= X {
		return "YES"
	}
	return "NO"
}

func main() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var N, X int
		fmt.Scanln(&N, &X)
		result := canHostParty(N, X)
		fmt.Println(result)
	}
}
