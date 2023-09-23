package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var a , b ,c int
	  fmt.Scan(&a , &b , &c)
	  
	  if a >= 10 && b >= 10 && c >= 10 && a+b+c >= 100  {
	    fmt.Println("PASS")
	  }else {
	    fmt.Println("FAIL")
	  }
	}
}
