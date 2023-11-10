package main
import "fmt"

func main(){
	var X ,Y int
	fmt.Scan(&X ,&Y)
	
	clearDays := 7 - X - Y 
	
	fmt.Println(clearDays)
}
