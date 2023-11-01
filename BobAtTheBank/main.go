package main
import "fmt"

func main(){
	var T , W ,X , Y , Z int
	fmt.Scan(&T)
	
	for i := 0; i < T; i++ {
	  fmt.Scan(&W ,&X ,&Y, &Z)
	  finalbalance := W + X*Z - Y*Z
	  fmt.Println(finalbalance)
	}
}
