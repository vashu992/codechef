package main
import "fmt"

func main(){
	var x , n , m int
	fmt.Scan(&x ,&n ,&m)
	
	if x+m >= n {
	  fmt.Println("YES")
	}else {
	  fmt.Println("NO")
	}
}
