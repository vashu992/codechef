package main
import "fmt"

func main(){
	var t int
	fmt. Scan(&t)
	
	for i := 0; i < t; i++ {
	  
	  var x int 
	  fmt.Scan(&x)
	  
	  if x<= 3 {
	    fmt.Println("BRONZE")
	  }else if x >6{
	    fmt.Println("GOLD")
	  }else  {
	    fmt.Println("SILVER")
	  }
	}
}
