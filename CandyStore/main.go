package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0;  i < t; i++ {
	  var x , y int
	  fmt.Scan(&x , &y)
	  
	  if y <= x {
	    fmt.Println(y)
	  }else {
	    earnings := x + (y - x)*2
	    fmt.Println(earnings)
	  }
	}
	
}
