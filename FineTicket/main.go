package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var x , p , q int
	  fmt.Scan(&x ,&p , &q)
	  
	  fine := x * (p -q)
	  fmt.Println (fine)
	}
}
