package main
import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	
	if x < 3 {
	  fmt.Println("GOLD")
	}else if x < 6 {
	  fmt.Println("SILVER")
	}else {
	  fmt.Println("BRONZE")
	}
}
