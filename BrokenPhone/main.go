package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var x , y int
	  fmt.Scan(&x , &y)
	  
	  if x < y {
	    fmt.Println("REPAIR")
	  }else if x > y {
	    fmt.Println("NEW PHONE")
	  }else {
	    fmt.Println("ANY")
	  }
	}
	
}
