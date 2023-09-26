package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  

	var n , m , k int
	fmt.Scan(&n , &m , &k)
	
	maxLoaves := m * k 
	
	if n <= maxLoaves {
	  fmt.Println("YES")
	}else {
	  fmt.Println("NO")
	}
}
	}
