package main
import "fmt"

func main(){
	var T int
	fmt.Scan(&T)
	
	for i := 0; i < T; i++ {
	  var X int
	  fmt.Scan(&X)
	  
	  var Spice string
	  if X < 4 {
	    Spice = "MILD"
	  }else if X < 7 {
	    Spice = "MEDIUM"
	  }else {
	    Spice = "HOT"
	  }
	  
	  fmt.Println(Spice)
	}
}
