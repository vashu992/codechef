package main
import "fmt"

func main(){
	var t int 
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  
	  var n , m , x int
	  fmt.Scan(&n , &m , &x)
	  
	  frameLenth := 2 *(n + m)
	  
	  totalCost := frameLenth * x 
	  
	  fmt.Println(totalCost)
	}
}
