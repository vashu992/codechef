package main
import "fmt"

func main(){
	var t int 
	fmt.Scan(&t)
	
	for i := 0; i < t; i ++ {
	  var x , y int
	  fmt.Scan(&x , &y)
	  
	  if y > x*5 {
	    fmt.Println("NO")
	  }else {
	    fmt.Println("YES")
	  }
	}
}
