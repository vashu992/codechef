package main
import "fmt"

func main(){
	var T int
	fmt.Scan(&T)
	
	for t := 0; t < T; t++ {
	  var A int
	  fmt.Scan(&A)
	  
	  isAlice := (A%2 == 0) && (A%7 == 0)
	  isBob := (A%2 == 1) && (A%9 == 0)
	  
	  if isAlice {
	    fmt.Println("Alice")
	  }else if isBob {
	    fmt.Println("Bob")
	  }else {
	    fmt.Println("Charlie")
	  }
	}
}
