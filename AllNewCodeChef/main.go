package main
import "fmt"

func main(){
	var X , Y int
	fmt.Scan(&X , &Y)
  
  if X < Y {
    fmt.Println("Old")
  }else if X > Y{
    fmt.Println("New")
  }else{
    fmt.Println("Same")
  }
}
