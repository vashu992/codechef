package main
import "fmt"

func main(){
	var responseCode int
	fmt.Scan(&responseCode)
	
	if responseCode == 404 {
	  fmt.Println("NOT FOUND")
	}else{
	  fmt.Println("FOUND")
	}
}
