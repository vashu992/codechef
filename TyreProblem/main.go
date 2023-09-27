package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var n , m int
	  fmt.Scan(&n , &m)
	  
	  total := (n * 2) + (m * 4)
	  
	  fmt.Println(total)
	}
}
