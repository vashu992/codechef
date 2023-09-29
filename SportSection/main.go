package main
import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	
	if x >= 8 && x <= 10 {
	  fmt.Println("YES")
	}else {
	  fmt.Println("NO")
	}
}
