package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	
	for i := 0; i < t; i++ {
	  var x , y ,z int
	  fmt.Scan(&x , &y , &z)
	  
	  totalPopulation := x - y + z
	  fmt.Println(totalPopulation)
	}
}
