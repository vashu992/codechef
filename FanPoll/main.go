package main
import "fmt"

func main(){
	var A , B, C int
	fmt.Scan(&A ,&B , &C) 
	
	if A > B && A > C {
	  fmt.Println("YES")
	}else {
	  fmt.Println("NO")
	}
}
