package main
import "fmt"

func main(){
	var n , u int 
	fmt.Scan(&n , &u)
	
	EligibleUsers := n - u 
	
	fmt.Println(EligibleUsers)
}
