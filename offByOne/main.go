package main
import "fmt"

func main(){
	var a , b int
	fmt.Scan(&a , &b)
	
	result := a + b
	
	result = result*10 + 1
	fmt.Println(result)
}
