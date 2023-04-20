package main

import "fmt"

func main() {
	t := 0
	fmt.Print("Enter number of test cases = ")
	fmt.Scanf("%d\n",&t)
	for t >0 {
		x ,y := 0,0
		
		fmt.Print("enter x ,y seperated by space =10 ")
		fmt.Scanf("%d\n",&x , x*y)
		t--
	}
}