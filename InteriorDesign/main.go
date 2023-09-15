package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var x1 , y1 , x2 , y2 int
	  fmt.Scan(&x1 , &y1 , &x2 , &y2)
	  
	  design1Cost := x1 + y1
	  design2Cost := x2 + y2
	  
	  if design1Cost < design2Cost {
	    fmt.Println(design1Cost)
	  }else {
	    fmt.Println(design2Cost)
	  }
	}
}
