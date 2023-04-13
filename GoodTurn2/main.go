package main

import "fmt"

func main() {
	testCase := 0
	fmt.Scanf("%d",testCase)

	for testCase > 0 {
		x , y := 0,0
		fmt.Scanf("%d",&x)
		fmt.Scanf("%d",&y)
		if x + y > 6 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
        testCase--
	}

	// for i := 0 ; i < testCase ; i++ {
	// 	x , y := 0,0
	// 	fmt.Scanf("%d",&x)
	// 	fmt.Scanf("%d",&y)
	// 	if x + y > 6 {
	// 		fmt.Println("YES")
	// 	}else{
	// 		fmt.Println("NO")
	// 	}
        
	}
