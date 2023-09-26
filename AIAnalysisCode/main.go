package main
import "fmt"

func main(){
	var c int
	fmt.Scan(&c)
	
	if c <= 1000 {
	  fmt.Println("YES")
	}else {
	  fmt.Println("NO")
	}
}
