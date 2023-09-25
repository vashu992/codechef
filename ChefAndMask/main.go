package main
import "fmt"

func main(){
	var t int 
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var x , y int
	  fmt.Scan(&x , &y)
	  
	  Disposable := x * 100
	  Cost := y * 10
	  
	  if Disposable < Cost {
	    fmt.Println("DISPOSABLE")
	  }else if Disposable > Cost {
	    fmt.Println("CLOTH")
	  }else {
	    fmt.Println("CLOTH")
	  }
	}
}
