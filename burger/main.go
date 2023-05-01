package main

import "fmt"

func main() {
	t := 0
	fmt.Print()
	fmt.Scanf("%d",&t)
	for t > 0 {
		a , b := 0, 0
		fmt.Print()
		fmt.Scanf("%d\n", &a)
		fmt.Print()
		fmt.Scanf("%d",&b)

		if a < b {
			fmt.Println(a)
		}else{
			fmt.Println(b)
		}
		t--
	}
}