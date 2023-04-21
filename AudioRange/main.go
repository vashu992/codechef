package main

import "fmt"

func main() {

	t := 0

	fmt.Print("")
	fmt.Scanf("%d\n",&t)

	for t > 0 {
		x := 0
		fmt.Print("")
		fmt.Scanf("%d\n",&x)

		if x >= 67 && x <= 45000 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
		t--
	}
}