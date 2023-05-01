package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d\n",&t)

	for i:= 1 ; i <= t ; i++ {
		x , y , a := 0,0,0
		fmt.Scan(&x,&y,&a )

		if a >= x && a < y {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}