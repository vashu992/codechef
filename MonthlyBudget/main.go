package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var x , y int
	  fmt.Scan(&x , &y)
	  
	  totalExpenditure := 30 *y 
	  if x >= totalExpenditure {
	    fmt.Println("YES")
	  }else {
	    fmt.Println("NO")
	  }
	}
}
