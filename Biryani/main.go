package main

import "fmt"

func main() {
	t := 0
	fmt.Print()
	fmt.Scanf("%d\n",&t)
	for t >0 {
		x ,y := 0,0
		
		fmt.Print()
		fmt.Scanf("%d %d",&x , &y)
		fmt.Println(x*y)
		t--
	}
}