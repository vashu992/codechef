package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var n, k int 
	  fmt. Scan(&n ,&k)
	  
	  BOBScore := n - k 
	  fmt.Println(BOBScore)
	}
}
