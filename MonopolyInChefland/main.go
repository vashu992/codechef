package main
import "fmt"

func main(){
	var T, R1, R2, R3 int
	fmt.Scan(&T)
	
	for i:=0; i <T; i++ {
	  fmt.Scan(&R1, &R2, &R3)
	  result := "NO"
	  if R1 > R2+R3 || R2 > R1+R3 || R3 > R1+R2 {
	    result = "YES"
	  }
	  fmt.Println(result)
	}
}
