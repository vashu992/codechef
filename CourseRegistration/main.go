package main
import "fmt"

func main(){
	var t , n , m , k  int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  //var n , m , k 
	  fmt.Scan(&n , &m , &k)
	  
	  total := n + k 
	  
	  if total <= m {
	    fmt.Println("YES")
	  }else {
	    fmt.Println("NO")
	  }
	}
}
