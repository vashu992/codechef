package main
import "fmt"

func main(){
	var t int 
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var k , x int
	  fmt.Scan(&k , &x)
	  
	  water := k - x 
	  fmt.Println(water)
	  
	}
}
